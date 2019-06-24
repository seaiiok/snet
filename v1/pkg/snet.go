package pkg

import (
	"errors"
	"fmt"
	"net"
	"snet/iface"
)

const (
	netWork = "tcp4"
)

type Server struct {
	Conf         map[string]string
	onConnect    func(*net.TCPConn)
	onDisConnect func(*net.TCPConn)
	onMessage    func(*net.TCPConn, iface.IPack)
	clients      map[string]client
}

type client struct {
	conn *net.TCPConn
}

func (s *Server) Start() {

	tcpAddr, err := net.ResolveTCPAddr(netWork, fmt.Sprintf("%s:%s", s.Conf["host"], s.Conf["port"]))
	if err != nil {
		fmt.Println("Server Start:", err)
		return
	}
	l, err := net.ListenTCP(netWork, tcpAddr)
	if err != nil {
		fmt.Println("Server Listen:", err)
		return
	}
	fmt.Println("Server ON:", s.Conf["host"], ":", s.Conf["port"])

	go func() {
		for {
			conn, err := l.AcceptTCP()
			if err != nil {
				fmt.Println("Server AcceptTcp:", err)
				continue
			} else {
				s.onConnect(conn)
				s.addClient(conn)
			}
			p := &Pack{}
			go s.onMessage(conn, p)
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) OnConnect(onConnect func(*net.TCPConn)) {
	s.onConnect = onConnect
}

func (s *Server) OnMessage(onMessage func(*net.TCPConn, iface.IPack)) {
	s.onMessage = onMessage
}

func (s *Server) OnDisConnect(onDisConnect func(*net.TCPConn)) {
	s.onDisConnect = onDisConnect
}

func onMessager(conn *net.TCPConn) {

}

func (s *Server) addClient(conn *net.TCPConn) {
	clientName := conn.RemoteAddr().String()
	newClient := client{
		conn: conn,
	}
	if _, found := s.clients[clientName]; found != true {
		s.clients[clientName] = newClient
		// return nil
	}
	// return errors.New("this client is exist,add client failed")
}

func (s *Server) deleteClient(clientName string) error {
	if _, found := s.clients[clientName]; found == true {
		delete(s.clients, clientName)
		return nil
	}
	return errors.New("this client is not exist,delete client failed")
}

type Pack struct {
}

func (p *Pack) Packet() {

}

func (p *Pack) UnPacket() {

}
