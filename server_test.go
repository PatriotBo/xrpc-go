package xrpc

import (
	"context"
	"fmt"
	"testing"
	"xrpc-go/proto/auth"
)

type Auth struct {
}

var mAuth = &Auth{}

func (a *Auth) Get(ctx context.Context, req *auth.AuthReq, resp *auth.AuthResp) error {
	uid := req.GetUid()
	name := req.GetName()

	fmt.Printf("uid = %d, name = %s \n", uid, name)
	return nil
}

type tfunc func(ctx context.Context, req *auth.AuthReq, resp *auth.AuthResp) error

func TestRegister(t *testing.T) {
	//server := NewServer(1, 1)
	//err := server.RegisterWithName(mAuth, "api-auth", &auth.AuthReq{}, &auth.AuthResp{})
	//if err != nil {
	//	fmt.Println("register failed err = ", err)
	//	return
	//}
	//service, ok := server.ServiceMap.Load("api-auth")
	//if ok {
	//	fmt.Println(service)
	req := new(auth.AuthReq)
	resp := new(auth.AuthResp)
	//	err := service.(*Service).Methods["Get"].Handler(service, context.Background(), req, resp)
	//	if err != nil {
	//		fmt.Println("ERROR = ", err)
	//	}
	//}
	var ti interface{}
	var tt tfunc
	tt = func(ctx context.Context, req *auth.AuthReq, resp *auth.AuthResp) error {
		fmt.Println(req)
		return nil
	}
	ti = tt
	ti.(MethodHandler)(mAuth, context.TODO(), req, resp)
}
