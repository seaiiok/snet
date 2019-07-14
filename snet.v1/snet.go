package snet

import (
	"fmt"
	"io"
	"net"
)

const (
	netWork = "tcp4"
)

type snet struct {
	onConnect     func(*Connection)
	onDisConnect  func(*Connection)
	onSendMessage func(*Connection)
	onRecvMessage func(*Connection, []byte)
}

//Connection 连接对象
type Connection struct {
	Conn *net.TCPConn
	Snet *snet
	p    Package
}

//New snet对象
func New(ip, port string) *snet {
	snet := new(snet)
	snet.start(ip, port)
	return snet
}

func (s *snet) start(ip, port string) {
	tcpAddr, err := net.ResolveTCPAddr(netWork, fmt.Sprintf("%s:%s", ip, port))
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

			buf := make([]byte, 1024)

			go func() {
				for {
					cnt, err := conn.Read(buf)
					if err != nil {
						if err == io.EOF {
							s.onDisConnect(newConn)
							return
						}
						// fmt.Println("Read:", err)
						s.onDisConnect(newConn)
						return
					}

					go s.onRecvMessage(newConn, newConn.UnPack(buf[:cnt]))
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

func (s *snet) OnRecvMessage(onRecvMessage func(conn *Connection, msg []byte)) {
	s.onRecvMessage = onRecvMessage
}

func (s *snet) OnSendMessage(onSendMessage func(conn *Connection)) {
	s.onSendMessage = onSendMessage
}

//OnSendMsg 发送消息
func (c *Connection) OnSendMsg(msg []byte) {
	_, err := c.Conn.Write(c.p.Pack(msg))
	if err != nil {
		// fmt.Println("Write:", err)
		c.Snet.onDisConnect(c)
		return
	}
}

//Pack 封包
func (c *Connection) Pack(msg []byte) []byte {
	return c.p.Pack(msg)
}

//UnPack 解包
func (c *Connection) UnPack(msg []byte) []byte {
	return c.p.UnPack(msg)
}
