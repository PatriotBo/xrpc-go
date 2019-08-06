package xrpc

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type requestI interface {
	Marshal() ([]byte, error)
}

type responseI interface {
	Unmarshal([]byte) error
}

<<<<<<< HEAD
type MethodHandler func(svr interface{}, ctx context.Context, req requestI, resp responseI) error

// 方法类型
type MethodType struct {
	Name    string // 完整方法名 服务名.方法名
	Alg     reflect.Value
	Reply   reflect.Value
	Handler MethodHandler
=======
type methodHandler func(svr interface{}, ctx context.Context, req requestI, resp responseI) error

// 方法类型
type MethodType struct {
	Name  string // 完整方法名 服务名.方法名
	Alg   reflect.Value
	Reply reflect.Value
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
}

type Service struct {
	Name    string
	Methods map[string]MethodType
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
func (s *server) RegisterWithName(svr interface{}, name string, arg, reply interface{}) error {
	if len(name) <= 0 {
		return ErrNoName
	}
	s.register(svr, name, arg, reply)
	return nil
}

func (s *server) register(sd interface{}, name string, arg, reply interface{}) {
	svr := &Service{
		Name:    name,
		Methods: make(map[string]MethodType),
	}

	sdType := reflect.TypeOf(sd)
	fmt.Println(sdType.NumMethod())
	for i := 0; i < sdType.NumMethod(); i++ {
		method := sdType.Method(i)
		fmt.Println("method: ", method.Name)
		me := MethodType{
<<<<<<< HEAD
			Name:    method.Name,
			Alg:     reflect.ValueOf(arg),
			Reply:   reflect.ValueOf(reply),
			Handler: method.Func.Interface().(MethodHandler),
=======
			Name:  method.Name,
			Alg:   reflect.ValueOf(arg),
			Reply: reflect.ValueOf(reply),
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
		}
		svr.Methods[method.Name] = me
	}

	s.ServiceMap.Store(name, svr) // 将服务注册进server
}
