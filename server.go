package xrpc

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"sync"
	"xrpc-go/proto/base"
)

type ProtoI interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type methodHandler func(svr interface{}, ctx context.Context, body []byte) (ProtoI, error)

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
	ServiceMap sync.Map          // 一个server可提供多个服务 map[string]*service
	serve      bool              // 服务是否开启
	conns      map[net.Conn]bool // 管理该server上的所有连接，stop时主动关闭所有连接
	sIp        string
	sPort      int
}

type Request struct {
	len int
	body []byte
}

type Response struct {
	len int
	body []byte
}

type connManager struct {
	conn net.Conn
	chanReq chan Request
	chanResp chan Response
}

func NewServer(ip string, port int) *server {
	return &server{
		sIp:   ip,
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

	// 将服务提供的方法注册进去
	for _, method := range sd.Methods {
		srv.methods[method.MethodName] = method
	}

	if s.serve {
		panic("found that register service while server is running")
	}

	if _, ok := s.ServiceMap.LoadOrStore(sd.ServiceName, srv); !ok {
		fmt.Println("register success")
	}
}

// 启动服务
func (s *server) Serve() error {
	addr := fmt.Sprintf("%s:%d", s.sIp, s.sPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("net.Listen failed: ", err)
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn accept failed: ", err)
			continue
		}
		s.conns[conn] = true
		// 处理连接
		cm := &connManager{
			conn: conn,
			chanResp: make(chan Response),
			chanReq: make(chan Request),
		}
		go s.handleConn(cm)
	}
}

func (s *server) handleConn(cm *connManager) {
	fmt.Println("conn established from remote addr: ", cm.conn.RemoteAddr().String())
	for {
		_, body, err := readRequest(cm.conn)
		if err != nil {
			fmt.Println("readRequest failed:")
			continue
		}
		req := Request{
			len: len(body),
			body:body,
		}
		// 开启goroutine 处理请求
		go s.handleRequest(cm)
		go s.handleResp(cm)

		cm.chanReq <- req
	}
}

// 解析请求数据包
func readRequest(c net.Conn) (header *RpcHead, body []byte, err error) {
	hb := make([]byte, RPC_HEAD_SIZE)
	_, err = io.ReadFull(c, hb)
	if err != nil {
		fmt.Println("read head failed: ", err)
		return
	}
	header, err = ParseHead(hb)
	if err != nil {
		return
	}
	body = make([]byte, header.Len)
	_, err = io.ReadFull(c, body)
	if err != nil {
		fmt.Println("read body failed:", err)
		return
	}
	return
}

// 从chanResp通道中取出响应数据 写入连接
func (s *server)handleResp(cm *connManager){
	for resp := range cm.chanResp {
		dlen := int32(len(resp.body))
		buf := new(bytes.Buffer)
		err := binary.Write(buf,binary.BigEndian,dlen)
		if err != nil {
			fmt.Println("handleResp write body len failed:",err)
			continue
		}

		err = binary.Write(buf,binary.BigEndian,resp.body)
		if err != nil {
			fmt.Println("handleResp write body failed:",err)
			continue
		}
		//fmt.Printf("handleResp ============= buf body:%+v \n",buf)
		wlen, err := cm.conn.Write(buf.Bytes())
		if wlen < int(dlen) || err != nil {
			fmt.Printf("write response into conn failed: wlen=%d, dlen=%d, err=%v\n",wlen,dlen,err)
			continue
		}
	}
}

// 从chanReq通道中取出请求数据， 进行处理，处理完成后添加到chanResp通道
func (s *server)handleRequest(cm *connManager) {
	for req := range cm.chanReq {
		s.serveRequest(cm,req.body)
	}
}

func (s *server) serveRequest(cm *connManager, body []byte) {
	rpcReq := new(pbBase.RpcReq)
	if err := rpcReq.Unmarshal(body); err != nil {
		fmt.Println("unmarshal socket body failed:", err)
		return
	}
	//fmt.Printf("serveRequest ---------------- rpcReq= %+v \n",rpcReq)
	svrName := rpcReq.GetRpc() // 获取服务名 api-auth.Get
	sn, mn, err := checkServiceName(svrName)
	if err != nil {
		return
	}
	if srv, ok := s.ServiceMap.Load(sn); ok {
		if method, ok := srv.(*service).methods[mn]; ok {
			resp, err := method.Handler(srv.(*service).server, context.TODO(), []byte(rpcReq.GetBody()))
			if err != nil {
				cm.writeResponse(svrName, []byte(err.Error()), pbBase.RpcCode_ERR,rpcReq.Seq)
				return
			}
			body, err := resp.Marshal() // 业务返回的包体
			cm.writeResponse(svrName, body, pbBase.RpcCode_SUCCESS,rpcReq.Seq)
			return
		}
		fmt.Printf("method %s not found in service %s\n", mn, sn)
		return
	}
}

func checkServiceName(name string) (sn, mn string, err error) {
	full := strings.Split(name, ".")
	if len(full) != 2 {
		fmt.Println("bad service name: ", name)
		err = errors.New(fmt.Sprintf("bad service name %s", name))
		return
	}
	sn = full[0]
	mn = full[1]
	return
}

func (cm *connManager)writeResponse(sname string, body []byte, code pbBase.RpcCode,seq int32) {
	rpcResp := new(pbBase.RpcResp)
	rpcResp.Body = string(body)
	rpcResp.Rpc = sname
	rpcResp.Code = int32(code)
	rpcResp.Seq = seq

	dAtA, err := rpcResp.Marshal()
	if err != nil {
		fmt.Println("rpcResp.Marshal failed while write response for service ", sname)
		//TODO:
	}
	//fmt.Printf("writeResponse for service %s with resp = %+v\n",sname,rpcResp)
	cm.chanResp <- Response{len:len(dAtA),body:dAtA}
}
