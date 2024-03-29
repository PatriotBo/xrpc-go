package xrpc

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
	"xrpc-go/proto/base"
)

const (
	RequestChanSize  = 16 // 每个连接的请求通道的缓冲数量 具体数值有待进一步确定
	ResponseChanSize = 16 // 每个连接的响应通道的缓冲数量 具体数值有待进一步确定
	RequestTimeout   = 3  // 请求超时时间 3s
)

type ClientConn struct {
	conn     net.Conn
	chanReq  chan request
	chanResp chan response
	count    int      // 连接中处理的请求数量 用于生成请求的序列号
	reqMap   sync.Map // map[seq]request 保存连接中正在处理的所有请求
}

type request struct {
	Seq       int    // 序列号 唯一标识连接上的一个请求
	Sname     string // 服务名
	Head      RpcHead
	Body      []byte
	ReplyChan chan response
}

type response struct {
	pbBase.RpcResp
}

/*
三种方案：
1. 客户端不维持长连接，每次调用服务时重新建立连接（DialTcp）
这样可以使得客户端结构简单清晰
2. 类似于（1），只不过对tcp连接做了缓存，类似与连接池的方案，但每次请求都要独占tcp连接（多个请求顺序执行，阻塞式）
3. 只建立一个长连接，所有请求共享一个连接，通过一些请求id来标识返回的响应数据属于哪个请求（需要超时机制，心跳检测机制，确定连接处于可用状态）
*/

// 一个客户端可以调用多个服务，因此，当服务不在同一个服务器上时，需要维护到各个不同服务器的多个长连接
type Client struct {
	connMap map[string]*ClientConn // map[addr]cc
	mu      sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		connMap: make(map[string]*ClientConn),
	}
}

func (c *Client) Call(ctx context.Context, sname string, req, resp ProtoI) error {
	addr, err := findServer(sname)
	if err != nil {
		fmt.Println("findServer failed:", err)
		return err
	}

	// 从client中获取连接信息，若没有，则创建连接
	c.mu.Lock()
	var reqData *request
	if cc, ok := c.connMap[addr]; ok { // 连接已存在
		cc.count++
		body, err := req.Marshal()
		if err != nil {
			fmt.Println("request marshal failed:", err)
			c.mu.Unlock()
			return err
		}
		reqData = &request{
			Head:      NewRpcHead(1, 0),
			Seq:       cc.count,
			Sname:     sname,
			Body:      body,
			ReplyChan: make(chan response),
		}
		if _, ok := cc.reqMap.LoadOrStore(reqData.Seq, reqData); ok {
			fmt.Printf("reqMap sotre failed request = %v exists! \n", reqData)
			return errors.New("request exists")
		}
		//fmt.Printf("connection exists call reqData:%+v \n",reqData)
		c.mu.Unlock()
		cc.chanReq <- *reqData
	} else { // 创建新连接
		clientConn, err := c.Connect(addr)
		if err != nil {
			fmt.Printf("new connect to addr[%s] failed:%v\n", addr, err)
			c.mu.Unlock()
			return err
		}
		c.connMap[addr] = clientConn
		c.mu.Unlock()

		clientConn.count++
		body, err := req.Marshal()
		if err != nil {
			fmt.Println("request marshal failed:", err)
			return err
		}
		reqData = &request{
			Head:      NewRpcHead(1, 0),
			Seq:       clientConn.count,
			Sname:     sname,
			Body:      body,
			ReplyChan: make(chan response,1),
		}

		if _, ok := clientConn.reqMap.LoadOrStore(reqData.Seq, reqData); ok {
			fmt.Printf("new conn reqMap sotre failed request = %v exists! \n", reqData)
			return errors.New("request exists")
		}
		//fmt.Println("new connection established  req: ",reqData)
		clientConn.chanReq <- *reqData
	}

	timer := time.Tick(time.Second * RequestTimeout)
	select {
	case <-timer:
		fmt.Println("request timeout")
		return errors.New("timeout")
	case respBody := <-reqData.ReplyChan:
		//fmt.Println("select response back =-=-=-=-=-=-=-=-=-=-=",string(respBody.Body) )
		if respBody.Rpc != sname {
			fmt.Printf("wrong response from [%s] for request service [%s]\n", respBody.Rpc, sname)
			return errors.New("wrong response")
		}
		err = resp.Unmarshal([]byte(respBody.Body))
		if err != nil {
			fmt.Printf("unmarshal resp failed for request service[%s]: %v\n", sname, err)
			return errors.New("unmarshal resp failed")
		}
	}
	return nil
}

//todo:根据服务发现机制，由服务名找到服务所在ip,port
func findServer(sname string) (string, error) {
	var ip string
	var port int
	ip = "localhost"
	port = 8090
	return fmt.Sprintf("%s:%d", ip, port), nil
}

// 创建新连接
func (c *Client) Connect(addr string) (*ClientConn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("dial tcp addr[%s] failed:%v \n", addr, err)
		return nil, err
	}

	cc := &ClientConn{
		conn:     conn,
		count:    0,
		chanReq:  make(chan request, RequestChanSize),
		chanResp: make(chan response, ResponseChanSize),
	}

	// 开启守护线程 处理该连接上的请求和响应
	go cc.sendRequest()
	go cc.recvResponse()
	return cc, nil
}

// 从连接中读取响应数据
func (cc *ClientConn) recvResponse() {
	for {
		// 首先读取4byte的长度信息
		var blen int32 // 头部指示的包体长度
		err := binary.Read(cc.conn, binary.BigEndian, &blen)
		if err != nil {
			fmt.Printf("recvResponse read header failed: blen = %d, err = %v\n", blen, err)
			continue
		}

		// 然后从连接中读取响应长度的数据
		var body = make([]byte, blen)
		rlen, err := cc.conn.Read(body) // rlen 实际读出的长度
		if rlen != int(blen) || err != nil {
			fmt.Printf("recvResponse read body failed: rlen = %d blen = %d, err = %v", rlen, blen, err)
			continue
		}
		resp := new(pbBase.RpcResp)
		err = resp.Unmarshal(body)
		if err != nil {
			fmt.Println("recvResponse unmarshal resp failed:", err)
			continue
		}

		if resp.Code != 0 {
			fmt.Printf("rpc call service[%s] failed return code[%d] \n", resp.Rpc, resp.Code)
		}

		if req, ok := cc.reqMap.Load(int(resp.Seq)); ok {
			rp := response{*resp}
			req.(*request).ReplyChan <- rp
			// 从请求map中删除该请求
			cc.reqMap.Delete(int(resp.Seq))
			cc.count--
		}
	}
}

// 处理连接上的请求： 接收通道中的数据，写入连接
func (cc *ClientConn) sendRequest() {
	for req := range cc.chanReq {
		req.Head.Len = uint16(len(req.Body))
		pbReq := new(pbBase.RpcReq)
		pbReq.Rpc = req.Sname
		pbReq.Body = string(req.Body)
		pbReq.Seq = int32(req.Seq)

		// 写入连接的包体body
		body, err := pbReq.Marshal()
		if err != nil {
			fmt.Printf("req[%v] marshal failed:%v\n", req, err)
			continue
		}
		// 依次将请求头部和包体写入到连接中
		buf := new(bytes.Buffer)
		err = binary.Write(buf, binary.BigEndian, req.Head.Flag)
		if err != nil {
			fmt.Printf("write head flag[%v] failed:%v\n", req.Head.Flag, err)
			continue
		}
		err = binary.Write(buf, binary.BigEndian, req.Head.Ver)
		if err != nil {
			fmt.Printf("write head ver[%v] failed:%v\n", req.Head.Ver, err)
			continue
		}
		err = binary.Write(buf, binary.BigEndian, uint16(len(body)))
		if err != nil {
			fmt.Printf("write head len[%v] failed:%v\n", uint16(len(body)), err)
			continue
		}
		err = binary.Write(buf, binary.BigEndian, body)
		if err != nil {
			fmt.Printf("write body[%v] failed:%v\n", body, err)
			continue
		}

		_, err = cc.conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println("write req into conn failed:", err)
			continue
		}
	}
}

//todo: 重连

// todo: 心跳检测
