package xrpc

import "net"

type ClientConn struct {
	conn net.Conn
}