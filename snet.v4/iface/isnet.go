package iface

import "net"

type ISnet interface {
	OnConnect(*net.TCPConn)
	OnDisConnect(*net.TCPConn, string)
	OnRecvMessage(*net.TCPConn, []byte)
	OnSendMessage(*net.TCPConn)
	OnLocalCommand(*net.TCPConn, []byte, []byte)
}
