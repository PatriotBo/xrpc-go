package xrpc

import (
	"context"
	"fmt"
	"testing"
	"xrpc-go/proto"
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

func TestRegister(t *testing.T) {
	server := NewServer(1, 1)
	err := server.RegisterWithName(mAuth, "api-auth", &auth.AuthReq{}, &auth.AuthResp{})
	if err != nil {
		fmt.Println("register failed err = ", err)
		return
	}
	service, ok := server.ServiceMap.Load("api-auth")
	if ok {
		fmt.Println(service)
		data := []byte("hello world")
		srv := service.(*Service)
		method := srv.Methods["Get"]
		alg := proto.UnMarshal(data, method.Alg.Interface())
		fmt.Printf("arg = %+v \n ", alg)
	}
}
