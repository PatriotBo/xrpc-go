package xrpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"sync"
	"xrpc-go/proto/base"
)

type protoI interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type methodHandler func(svr interface{}, ctx context.Context, req interface{},body []byte)(protoI, error)

type MethodDesc struct {
	MethodName string
	Handler    methodHandler
}

type ServiceDesc struct {
	ServiceName string
	HandlerType interface{}
	Methods     []MethodDesc
}

type service struct {
	Name    string
	server  interface{}
	methods map[string]MethodDesc
}

type server struct {
	ServiceMap sync.Map // 一个server可提供多个服务 map[string]*service
	serve      bool     // 服务是否开启
	conns      map[net.Conn]bool // 管理该server上的所有连接，stop时主动关闭所有连接
	sIp        string
	sPort      int
}

func NewServer(ip string, port int) *server {
	return &server{
		sIp:ip,
		sPort: port,
		conns: make(map[net.Conn]bool),
	}
}

var (
	ErrNoName = errors.New("service name is necessary")
)

// 注册服务
func (s *server) RegisterService(sd *ServiceDesc, ss interface{}) {
	handleType := reflect.TypeOf(sd.HandlerType).Elem()
	serviceType := reflect.TypeOf(ss)
	if !serviceType.Implements(handleType) {
		panic(fmt.Sprintf("found the type %v does not satisfy %v", serviceType, handleType))
	}
	s.register(sd, ss)
}

func (s *server) register(sd *ServiceDesc, ss interface{}) {
	srv := &service{
		Name:    sd.ServiceName,
		server:  ss,
		methods: make(map[string]MethodDesc),
	}

	if s.serve {
		panic("found that register service while server is running")
	}

	if _, ok := s.ServiceMap.LoadOrStore(sd.ServiceName, srv); !ok {
		fmt.Println("register success")
	}
}

// 启动服务
func (s *server)Serve()error{
	addr := fmt.Sprintf("%s:%d",s.sIp, s.sPort)
	listener, err := net.Listen("tcp",addr)
	if err != nil {
		fmt.Println("net.Listen failed: ",err)
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn accept failed: ",err)
			continue
		}
		s.conns[conn] = true
		// 处理连接
		go s.handleConn(conn)
	}
}

func (s *server)handleConn(c net.Conn){
	fmt.Println("conn established from remote addr: ",c.RemoteAddr().String())
	for {
		_, body, err := readRequest(c)
		if err != nil {
			fmt.Println("readRequest failed:")
			continue
		}

		// 开启goroutine 处理请求
		go s.handleRequest(c, body)
	}
}

// 解析请求数据包
func readRequest(c net.Conn)(header *RpcHead,body []byte,err error){
	hb := make([]byte,RPC_HEAD_SIZE)
	_, err = io.ReadFull(c,hb)
	if err != nil {
		fmt.Println("read head failed: ",err)
		return
	}

	header, err = ParseHead(hb)
	if err != nil {
		return
	}

	body = make([]byte,header.Len)
	_, err = io.ReadFull(c,body)
	if err != nil {
		fmt.Println("read body failed:",err)
		return
	}
	return
}

func (s *server)handleRequest(c net.Conn,body []byte){
	rpcReq := new(pbBase.RpcReq)
	if err := rpcReq.Unmarshal(body); err != nil {
		fmt.Println("unmarshal socket body failed:",err)
		return
	}

	svrName := rpcReq.GetRpc() // 获取服务名 api-auth.Get
	sn, mn, err := checkServiceName(svrName)
	if err != nil {
		return
	}
	if srv, ok := s.ServiceMap.Load(sn);ok {
		if method,ok := srv.(service).methods[mn]; ok {

			resp, err := method.Handler(srv.(service).server,context.TODO(),rpcReq,[]byte(rpcReq.GetBody()))
			if err != nil {

			}
			body,err := resp.Marshal()

		}
		fmt.Printf("method %s not found in service %s\n",mn,sn)
		return
	}

}

func checkServiceName(name string)(sn,mn string,err error){
	full := strings.Split(name,".")
	if len(full) != 2 {
		fmt.Println("bad service name: ",name)
		err = errors.New(fmt.Sprintf("bad service name %s",name))
		return
	}
	sn = full[0]
	mn = full[1]
	return
}