package iface

import "net"

type IServer interface {
	Start()
	Stop()
	OnConnect(func(*net.TCPConn))
	OnMessage(func(*net.TCPConn, IPack))
	OnDisConnect(func(*net.TCPConn))
}

type IPack interface {
	Packet()
	UnPacket()
}

// type IClient interface {
// 	ClientStart()
// 	ClientStop()
// }

// type IConnections interface {
// 	AddClient(*net.TCPConn) error
// 	DeleteClient(string) error
// }
