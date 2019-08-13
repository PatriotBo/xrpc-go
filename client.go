package xrpc

import (
	"net"
	"sync"
)

type ClientConn struct {
	conn     net.Conn
	chanReq  request
	chanResp Response
	mu       sync.Mutex
}

type request struct {
	Seq       int // 序列号 唯一标识连接上的一个请求
	replyChan chan interface{}
}

/*
三种方案：
1. 客户端不维持长连接，每次调用服务时重新建立连接（DialTcp）
这样可以使得客户端结构简单清晰
2. 类似于（1），只不过对tcp连接做了缓存，类似与连接池的方案，但每次请求都要独占tcp连接（多个请求顺序执行，阻塞式）
3. 只建立一个长连接，所有请求共享一个连接，通过一些请求id来标识返回的响应数据属于哪个请求（需要超时机制，心跳检测机制，确定连接处于可用状态）
*/
type Client struct {
	connMap map[string]ClientConn // map[addr]cc
}
