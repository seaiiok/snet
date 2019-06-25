package snet

import (
	"fmt"
	"net"
)

const (
	iP      = "127.0.0.1"
	port    = "494"
	netWork = "tcp4"
)

type Snet struct {
	conf         string
	onConnect    func(*net.TCPConn)
	onDisConnect func(*net.TCPConn)
	onMessage    func(*net.TCPConn, string)
}

func New(addr string) *Snet {
	snet := &Snet{
		conf: "ip+port",
	}
	snet.start()
	return snet
}

func (s *Snet) start() {
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
	fmt.Println("Server ON:", iP, ":", port)

	go func() {
		for {
			conn, err := l.AcceptTCP()
			if err != nil {
				fmt.Println("Server AcceptTcp:", err)
				continue
			}
			fmt.Println(conn.RemoteAddr().String())
			s.onConnect(conn)

			buf := make([]byte, 512)

			go func() {
				for {
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println(err)
						s.onDisConnect(conn)
						return
					}
					s.onMessage(conn, string(buf[:cnt]))

					_, err = conn.Write(buf[:cnt])
					if err != nil {
						fmt.Println(err)
						s.onDisConnect(conn)
						return
					}

				}

			}()

		}
	}()
}

func (s *Snet) OnConnect(onConnect func(conn *net.TCPConn)) {
	s.onConnect = onConnect
}

func (s *Snet) OnDisConnect(onDisConnect func(conn *net.TCPConn)) {
	s.onDisConnect = onDisConnect
}

func (s *Snet) OnMessage(onMessage func(conn *net.TCPConn, msg string)) {
	s.onMessage = onMessage
}
