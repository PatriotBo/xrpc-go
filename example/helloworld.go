package main

import (
	"context"
	"fmt"
	"xrpc-go/proto/hello"
)

type Hello struct {
}

var mHelloWorld = &Hello{}

func (h *Hello)SayHello(ctx context.Context,req pbHello.HelloReq)(*pbHello.HelloResp,error){
	fmt.Printf("********HelloWorld******** req:%v\n ",req)

	uid := req.GetUid()
	msg := req.GetMsg()

	fmt.Printf("user[%d] says `%s` to the world \n",uid,msg)
	resp := new(pbHello.HelloResp)
	resp.Code = 0
	resp.Msg = "success"
	resp.Info = fmt.Sprintf("hello %d",uid)
	return resp,nil
}