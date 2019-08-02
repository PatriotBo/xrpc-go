package xrpc

import (
	"context"
	"sync"
)

type requestI interface {
	Marshal() ([]byte, error)
}

type responseI interface {
	UnMarshal([]byte) error
}

type methodHandler func(ctx context.Context, req requestI, resp responseI)error

// 方法类型
type MethodType struct {
	Name string // 完整方法名 服务名.方法名
	Handler methodHandler
}

type service struct {
	Name  string
	SId   int32
	SType int32
	Methods map[string]MethodType
}

type Request struct {
	Sname  string	// 服务名
	Mname  string   // 方法名
	Body  []byte    // 请求内容 参数等信息

	//lock sync.Mutex
}

type Response struct {

}

type server struct {
	ServiceMap  sync.Map // 一个server可提供多个服务 map[string]*service
	Req Request
	Resp Response
}