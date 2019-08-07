package xrpc

import (
	"context"
	"fmt"
	"testing"
	"xrpc-go/proto/auth"
)

type AuthServer interface {
	Get(ctx context.Context, req auth.AuthReq) (*auth.AuthResp,error)
}

func _Get_Handler(svr interface{}, ctx context.Context, req interface{},body []byte)(protoI, error) {
	fmt.Println("_Get_Handler")
	req.(protoI).Unmarshal(body)
	return svr.(AuthServer).Get(ctx, req.(auth.AuthReq))
}

type Auth struct {
}

var mAuth = &Auth{}

func (a *Auth) Get(ctx context.Context, req auth.AuthReq)(*auth.AuthResp, error) {
	uid := req.GetUid()
	name := req.GetName()
	resp := new(auth.AuthResp)
	fmt.Printf("uid = %d, name = %s \n", uid, name)
	return resp, nil
}

func TestRegister(t *testing.T) {
	server := NewServer("127.0.0.1", 8090)
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
		//s.(service).methods["Get"].Handler()
		resp, err := s.(AuthServer).Get(context.TODO(), *req)
		if err != nil {
			fmt.Println("call service api-auth.Get failed: ", err,resp)
		}
	}
}
