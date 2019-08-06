package xrpc

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

type requestI interface {
	Marshal() ([]byte, error)
}

type responseI interface {
	Unmarshal([]byte) error
}

type methodHandler func(svr interface{}, ctx context.Context, req interface{}, resp unsafe.Pointer) error

// 方法类型
type MethodType struct {
	Name  string // 完整方法名 服务名.方法名
	Alg   reflect.Value
	Reply reflect.Value
}

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

type Request struct {
	Sname string // 服务名
	Mname string // 方法名
	Body  []byte // 请求内容 参数等信息

	//lock sync.Mutex
}

type Response struct {
	Sname string // 服务名
	Mname string // 方法名
	Body  []byte // 响应内容
}

type server struct {
	ServiceMap sync.Map // 一个server可提供多个服务 map[string]*service
	serve      bool     // 服务是否开启
	SId        int32
	SType      int32
}

func NewServer(sid, stype int32) *server {
	return &server{
		SId:   sid,
		SType: stype,
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
