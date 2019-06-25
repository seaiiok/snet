package snet

import (
	"fmt"
	"io"
	"net"
)

const (
	iP      = "127.0.0.1"
	port    = "495"
	netWork = "tcp4"
)

type snet struct {
	onConnect     func(*Connection)
	onDisConnect  func(*Connection)
	onSendMessage func(*Connection)
	onRecvMessage func(*Connection, Package)
}

type Connection struct {
	Conn *net.TCPConn
	Snet *snet
}

func New() *snet {
	snet := new(snet)
	snet.start()
	return snet
}

func (s *snet) start() {
	tcpAddr, err := net.ResolveTCPAddr(netWork, fmt.Sprintf("%s:%s", iP, port))
	if err != nil {
		fmt.Println("Server Start:", err)
		return
	}
	l, err := net.ListenTCP(netWork, tcpAddr)
	if err != nil {
		fmt.Println("Server Listen:", err)
		return
	}

	go func() {
		for {
			conn, err := l.AcceptTCP()
			if err != nil {
				fmt.Println("Server AcceptTcp:", err)
				continue
			}
			newConn := new(Connection)
			newConn.Conn = conn
			newConn.Snet = s

			s.onConnect(newConn)

			buf := make([]byte, 512)
			pge := Package{}

			go func() {
				for {
					cnt, err := conn.Read(buf)
					if err != nil {
						if err == io.EOF {
							// fmt.Println("Read:",err)
							return
						}
						fmt.Println("Read:", err)
						s.onDisConnect(newConn)
						return
					}

					go s.onRecvMessage(newConn, pge.UnPack(buf[:cnt]))
					go s.onSendMessage(newConn)
				}

			}()

		}
	}()
}

func (s *snet) OnConnect(onConnect func(conn *Connection)) {
	s.onConnect = onConnect
}

func (s *snet) OnDisConnect(onDisConnect func(conn *Connection)) {
	s.onDisConnect = onDisConnect
}

func (s *snet) OnRecvMessage(onRecvMessage func(conn *Connection, msg Package)) {
	s.onRecvMessage = onRecvMessage
}

func (s *snet) OnSendMessage(onSendMessage func(conn *Connection)) {
	s.onSendMessage = onSendMessage
}

func (c *Connection) OnSendMsg(msg Package) {
	_, err := c.Conn.Write(msg.Pack())
	if err != nil {
		fmt.Println("Write:", err)
		c.Snet.onDisConnect(c)
		return
	}
}
