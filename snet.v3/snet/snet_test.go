package snet

import (
	"fmt"
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

// func BenchmarkSnet(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		clientGo()
// 	}
// }

type Client struct {
}

func serverGo() {
	c := &Client{}
	s := NewServer("localhost", 496)
	s.AddConnection(c)
	s.Serve()
}

func (c *Client) OnConnect(conn *net.TCPConn) {
	gcmd.Println(gcmd.Info, "客户端连接:", conn.RemoteAddr())
}

func (c *Client) OnDisConnect(conn *net.TCPConn, reason string) {
	gcmd.Println(gcmd.Info, "客户端断开连接:", conn.RemoteAddr(), "原因:", reason)
}

func (c *Client) OnRecvMessage(conn *net.TCPConn, msg []byte) {
	gcmd.Println(gcmd.Ok, "接收客户端消息:", msg)
}

func (c *Client) OnSendMessage(conn *net.TCPConn) {

}

// //客户端
func clientGo() {
	conn, err := net.Dial("tcp", "127.0.0.1:496")
	if err != nil {
		gcmd.Println(gcmd.Warn, "client dial err, exit!")
		return
	}

	time.Sleep(1 * time.Second)

	p := &Package{}
	p.Msg = []byte{1, 2, 3}
	x := p.Pack()
	fmt.Println("x:", x)

	b := []byte{0, 0, 0, 3, 184, 77, 142, 166, 1, 2, 3, 4, 5, 6, 0, 0, 0, 3, 184, 77, 142, 166, 7, 8, 9, 10, 11, 12}

	// b := p.Pack()
	// gcmd.Println(gcmd.Ok, b)
	conn.Write(b)
	conn.Write(b)
	// for i := 0; i < 1; i++ {
	// 	p.Msg = append(p.Msg, []byte(fmt.Sprintf("%d", i))...)
	// }
	// b = p.Pack()

	// conn.Write(b)

}
