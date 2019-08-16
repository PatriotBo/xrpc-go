package xrpc

import (
	"context"
	"fmt"
	"testing"
	"xrpc-go/proto/auth"
)

type AuthServer interface {
	Get(ctx context.Context, req auth.AuthReq) (*auth.AuthResp, error)
}

func _Get_Handler(svr interface{}, ctx context.Context, body []byte) (ProtoI, error) {
	fmt.Println("_Get_Handler")
	aReq := new(auth.AuthReq)
	err := aReq.Unmarshal(body)
	if err != nil {
		fmt.Println("auth.AuthReq unmarshal failed:", err)
		return nil, err
	}
	return svr.(AuthServer).Get(ctx, *aReq)
}

type Auth struct {
}

var mAuth = &Auth{}

func (a *Auth) Get(ctx context.Context, req auth.AuthReq) (*auth.AuthResp, error) {
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

	server.Serve()
	//if s, ok := server.ServiceMap.Load("api-auth"); ok {
	//	fmt.Println("-------------------")
	//	req := new(auth.AuthReq)
	//	//s.(service).methods["Get"].Handler()
	//	resp, err := s.(AuthServer).Get(context.TODO(), *req)
	//	if err != nil {
	//		fmt.Println("call service api-auth.Get failed: ", err, resp)
	//	}
	//}
}
