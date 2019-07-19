package snet

import (
	"fmt"
	"gcom/gcmd"
	"io"
	"net"

	"snet/snet.v3/iface"
)

type Snet struct {
	IP   string
	Port int
	Conn iface.IConnection
}

func NewServer(ip string, port int) iface.ISnet {
	return &Snet{
		IP:   ip,
		Port: port,
	}
}

func (s *Snet) Serve() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		gcmd.Println(gcmd.Err, "server tcp addr err:", err)
		return
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		gcmd.Println(gcmd.Err, "server listen tcp err:", err)
		return
	}

	go func() {
		for {
			conn, err := l.AcceptTCP()
			if err != nil {
				gcmd.Println(gcmd.Err, "server accept tcp err:", err)
				// TODO 客户端断开连接
				if err == io.EOF {
					s.Conn.OnDisConnect(conn, "client close connection")
					continue
				}
				s.Conn.OnDisConnect(conn, err.Error())
				continue
			}
			// TODO 客户端连接
			s.Conn.OnConnect(conn)

			// TODO 连接处理协程
			go s.NewConnection(conn)
		}
	}()
}

func (s *Snet) Stop() {

}

func (s *Snet) AddConnection(connection iface.IConnection) {
	s.Conn = connection
}
