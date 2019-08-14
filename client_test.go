package xrpc

import (
	"context"
	"fmt"
	"testing"
	"xrpc-go/proto/auth"
)

func TestClient(t *testing.T) {
	client := NewClient()
	client.Connect("localhost:8090")

	sname := "api-auth.Get"
	req := &auth.AuthReq{
		Uid:2304,
		Name:"vincent",
	}
	resp := new(auth.AuthResp)
	err := client.Call(context.TODO(),sname,req,resp)
	if err != nil {
		fmt.Println("call failed:",err)
		return
	}
	fmt.Printf("call success resp:%v\n",*resp)
}