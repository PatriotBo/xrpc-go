package example

import (
	"context"
	"fmt"
	"xrpc-go/proto/hello"
)

type hello struct {
}

var HelloWorld = &hello{}

func (h *hello)HelloWorld(ctx context.Context,req *pbHello.HelloReq,resp *pbHello.RpcResp)error {
	fmt.Printf("********HelloWorld******** req:%v\n ",req)

	uid := req.GetUid()
	msg := req.GetMsg()

	fmt.Printf("user[%d] says `%s` to the world \n",uid,msg)
	resp.Code = 0
	resp.Msg = "success"
	resp.Info = fmt.Sprintf("hello %d",uid)
	return nil
}