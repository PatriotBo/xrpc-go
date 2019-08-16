package xrpc

import (
	"context"
	"fmt"
	"testing"
	"xrpc-go/proto/hello"
)

//var rpcClient *Client

func Init(){
	//rpcClient := NewClient()
	//rpcClient.Connect("localhost:8090")
}

func TestClient(t *testing.T) {
	rpcClient := NewClient()
	rpcClient.Connect("localhost:8090")

	sname := "api-hello.SayHello"
	req := &pbHello.HelloReq{
		Uid: 2304,
		Msg: "hello to you",
	}
	resp := new(pbHello.HelloResp)
	err := rpcClient.Call(context.TODO(),sname,req,resp)
	if err != nil {
		fmt.Println("call failed:",err)
		return
	}
	fmt.Printf("call success resp:%v\n",*resp)

}

func BenchmarkClient_Call(b *testing.B) {
	b.StopTimer()
	rpcClient := NewClient()
	rpcClient.Connect("localhost:8090")
	fmt.Println("benchmark time = ",b.N)
	b.StartTimer()
	for i := 0;i < b.N; i++ {
		sname := "api-hello.SayHello"
		req := &pbHello.HelloReq{
			Uid: 2304,
			Msg: "hello to you",
		}
		resp := new(pbHello.HelloResp)
		err := rpcClient.Call(context.TODO(),sname,req,resp)
		if err != nil {
			fmt.Println("call failed:",err)
			return
		}
		//fmt.Printf("call success resp:%v\n",*resp)
	}
}