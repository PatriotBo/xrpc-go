package xrpc

import (
	"context"
	"fmt"
	"testing"
	"unsafe"
	"xrpc-go/proto/auth"
)

type AuthServer interface {
	Get(ctx context.Context, req auth.AuthReq, resp *auth.AuthResp) error
}

func _Get_Handler(svr interface{}, ctx context.Context, req interface{}, resp unsafe.Pointer) error {
	fmt.Println("_Get_Handler")
	return svr.(AuthServer).Get(ctx, req.(auth.AuthReq), (*auth.AuthResp)(resp))
}

type Auth struct {
}

var mAuth = &Auth{}

func (a *Auth) Get(ctx context.Context, req auth.AuthReq, resp *auth.AuthResp) error {
	uid := req.GetUid()
	name := req.GetName()

	fmt.Printf("uid = %d, name = %s \n", uid, name)
	return nil
}

func TestRegister(t *testing.T) {
	server := NewServer(1, 1)
	sd := &ServiceDesc{
		ServiceName: "api-auth",
		HandlerType: (*AuthServer)(nil),
		Methods: []MethodDesc{
			{
				MethodName: "Get",
				Handler:    _Get_Handler,
			},
		},
	}
	server.RegisterService(sd, mAuth)

	if s, ok := server.ServiceMap.Load("api-auth"); ok {
		fmt.Println("-------------------")
		req := new(auth.AuthReq)
		resp := new(auth.AuthResp)
		err := s.(AuthServer).Get(context.TODO(), *req, resp)
		if err != nil {
			fmt.Println("call service api-auth.Get failed: ", err)
		}
	}
}
