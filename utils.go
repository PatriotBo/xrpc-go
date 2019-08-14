package xrpc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
)

const(
	RPC_HEAD_SIZE = 7 // 包头长度
)

type RpcHead struct {
	Flag [3]byte // "RPC"标记
	Ver  uint16  // 版本
	Len  uint16  // 包体长度
}

// 读取并校验包头
func ParseHead(buf []byte)(*RpcHead,error){
	var head RpcHead
	bufReader := bytes.NewBuffer(buf)
	if err := binary.Read(bufReader,binary.BigEndian,&head.Flag); err != nil {
		fmt.Println("read head.Flag failed:",err)
		return nil,err
	}

	if err := binary.Read(bufReader,binary.BigEndian,&head.Ver);err != nil {
		fmt.Println("read head.Ver failed:",err)
		return nil,err
	}

	if err := binary.Read(bufReader,binary.BigEndian,&head.Len);err != nil {
		fmt.Println("rad head.Len failed:",err)
		return nil, err
	}

	if head.Flag[0] != uint8('R') || head.Flag[1] != uint8('P') || head.Flag[2] != uint8('C') || head.Len <= 0 {
		fmt.Println("request header illegal !!!")
		return nil, errors.New("request header illegal !!!")
	}
	return &head,nil
}

// 解析包体中的服务数据

func NewRpcHead(ver, len uint16)RpcHead{
	var flag [3]byte
	flag[0] = uint8('R')
	flag[1] = uint8('P')
	flag[2] = uint8('C')
	return RpcHead{
		Flag: flag,
		Ver: ver,
		Len: len,
	}
}