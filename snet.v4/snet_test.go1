package snet

import (
	"gcom/gcmd"
	"net"
	"testing"
	"time"
)

func TestSnet(t *testing.T) {

	//启动Tcp服务器
	go serverGo()

	time.Sleep(1 * time.Second)

	//启动客户端
	go clientGo()

	select {}
}

type server struct {
}

func serverGo() {
	s := NewServer("localhost", "496", &server{})
	s.Serve()

}

func (this *server) OnConnect(conn *net.TCPConn) {
	gcmd.Println(gcmd.Info, "客户端连接:", conn.RemoteAddr())
}

func (this *server) OnDisConnect(conn *net.TCPConn, reason string) {
	gcmd.Println(gcmd.Info, "客户端断开连接:", conn.RemoteAddr(), "原因:", reason)
}

func (this *server) OnRecvMessage(conn *net.TCPConn, msg []byte) {
	gcmd.Println(gcmd.Ok, "接收客户端消息:", string(msg))
}

func (this *server) OnSendMessage(conn *net.TCPConn) {
	p := &Package{}

	msg := []byte("its a test msg!")
	packMsg, _ := p.Pack(msg)

	conn.Write(packMsg)
}

type client struct{}

func (this *client) OnConnect(conn net.Conn) {

}

func (this *client) OnDisConnect(conn net.Conn, reason string) {

}

func (this *client) OnRecvMessage(conn net.Conn, msg []byte) {

}

func (this *client) OnSendMessage(conn net.Conn) {

}

// //客户端
func clientGo() {
clients.

}
