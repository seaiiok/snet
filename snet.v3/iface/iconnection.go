package iface

import "net"

type IConnection interface {
	OnConnect(*net.TCPConn)
	OnDisConnect(*net.TCPConn, string)
	OnRecvMessage(*net.TCPConn, []byte)
	OnSendMessage(*net.TCPConn)
}
