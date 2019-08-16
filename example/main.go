package main

import (
	"context"
	"fmt"
	"xrpc-go"
	"xrpc-go/proto/hello"
)

type HelloWorld interface {
	SayHello(ctx context.Context,req pbHello.HelloReq)(*pbHello.HelloResp,error)
}

func _SayHello_Handler(svr interface{},ctx context.Context,body []byte)(xrpc.ProtoI,error) {
	req := new(pbHello.HelloReq)
	err := req.Unmarshal(body)
	if err != nil {
		fmt.Println("unmarshal request failed: ",err)
		return nil,err
	}
	return svr.(HelloWorld).SayHello(ctx, *req)
}

var controller = mHelloWorld

func main(){
	sr := xrpc.NewServer("127.0.0.1",8090)
	sd := &xrpc.ServiceDesc{
		ServiceName: "api-hello",
		HandlerType: (*HelloWorld)(nil),
		Methods: []xrpc.MethodDesc{
			{
				MethodName:"SayHello",
				Handler: _SayHello_Handler,
			},
		},
	}

	sr.RegisterService(sd, controller)

	err := sr.Serve()
	if err != nil {
		fmt.Println("server serve failed: ",err)
		return
	}
}