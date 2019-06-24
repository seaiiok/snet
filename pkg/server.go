package pkg

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"snet/iface"
// )

// type Server struct {
// 	netWork     string
// 	iP          string
// 	port        int
// 	contx       context.Context
// 	contxCancel context.CancelFunc
// }

// func (s *Server) Start() {

// 	tcpAddr, err := net.ResolveTCPAddr(s.netWork, fmt.Sprintf("%s:%d", s.iP, s.port))
// 	if err != nil {
// 		fmt.Println("Server Start:", err)
// 		return
// 	}
// 	l, err := net.ListenTCP(s.netWork, tcpAddr)
// 	if err != nil {
// 		fmt.Println("Server Listen:", err)
// 		return
// 	}
// 	fmt.Println("Server ON:", s.iP, ":", s.port)

// 	newClients := newClients(s.contx)
// 	go newClients.Start()

// 	go func() {
// 		for {
// 			conn, err := l.AcceptTCP()
// 			if err != nil {
// 				fmt.Println("Server AcceptTcp:", err)
// 			}

// 			select {
// 			case <-s.ServerContext.Contx.Done():
// 				return
// 			default:
// 				newClients.AddClient(conn)
// 			}
// 		}
// 	}()

// }

// func (s *Server) Stop() {

// }

// func NewServer() iface.IServer {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	return &Server{
// 		NetWork: "tcp4",
// 		IP:      "127.0.0.1",
// 		Port:    777,
// 		ServerContext: &Context{
// 			Contx:       ctx,
// 			ContxCancel: cancel,
// 		},
// 	}
// }
// const (
// 	netWork = "tcp4"
// )

// type Server struct {
// 	Conf map[string]string
// }

// func (s *Server) ServerStart() {
// 	tcpAddr, err := net.ResolveTCPAddr(netWork, fmt.Sprintf("%s:%d", Config.Host, Config.port))
// 	if err != nil {
// 		fmt.Println("Server Start:", err)
// 		return
// 	}
// 	l, err := net.ListenTCP(netWork, tcpAddr)
// 	if err != nil {
// 		fmt.Println("Server Listen:", err)
// 		return
// 	}
// 	fmt.Println("Server ON:", Config.Host, ":", Config.Port)

// 	newClients := newClients(s.contx)
// 	go newClients.Start()

// 	go func() {
// 		for {
// 			conn, err := l.AcceptTCP()
// 			if err != nil {
// 				fmt.Println("Server AcceptTcp:", err)
// 			}

// 			select {
// 			case <-s.ServerContext.Contx.Done():
// 				return
// 			default:
// 				newClients.AddClient(conn)
// 			}
// 		}
// 	}()
// }

// func (s *Server) ServerStart() {
// 	tcpAddr, err := net.ResolveTCPAddr(netWork, fmt.Sprintf("%s:%s", s.Conf["host"], s.Conf["port"]))
// 	if err != nil {
// 		fmt.Println("Server Start:", err)
// 		return
// 	}
// 	l, err := net.ListenTCP(netWork, tcpAddr)
// 	if err != nil {
// 		fmt.Println("Server Listen:", err)
// 		return
// 	}
// 	fmt.Println("Server ON:", s.Conf["host"], ":", s.Conf["port"])

// 	go func() {
// 		for {
// 			conn, err := l.AcceptTCP()
// 			if err != nil {
// 				fmt.Println("Server AcceptTcp:", err)
// 			}
// 			fmt.Println(conn.RemoteAddr().String())
// 		}
// 	}()
// }

// func (s *Server) ServerStop() {

// }

// // ------------------------------------------------------------------
// type Connections struct {
// }

// type Connection struct {
// }

// func (c *Connections) AddClient(conn *net.TCPConn) error {
// 	return nil
// }

// func (c *Connections) DeleteClient(clientKey string) error {
// 	return nil
// }

// func NewConnections() iface.IConnections {
// 	return &Connections{}
// }
