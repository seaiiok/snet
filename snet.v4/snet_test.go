package snet

import (
	"collector/api"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/seaiiok/gcom/gcmd"
	"snet/snet.v4/clients"
	"snet/snet.v4/packet"
)

func TestSnet(t *testing.T) {

	//启动Tcp服务器
	go serverGo()

	time.Sleep(1 * time.Second)

	//启动客户端
	// go clientGo()

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
	p := &packet.Package{}
	pm := &api.Msg{
		Cmd:  100,
		File: "c:/faf",
		Md5:  "fjisod8902309",
		Msg:  []byte("its a test msg!"),
	}

	b, _ := proto.Marshal(pm)

	packMsg, _ := p.Pack(b)

	conn.Write(packMsg)
}

//client ...
type client struct{}

func (this *client) OnConnect(conn net.Conn) {
	gcmd.Println(gcmd.Info, "连接到服务器:", conn.RemoteAddr())
}

func (this *client) OnDisConnect(conn net.Conn, reason string) {
	gcmd.Println(gcmd.Info, "断开服务器连接:", conn.RemoteAddr())
}

func (this *client) OnRecvMessage(conn net.Conn, msg []byte) {
	gcmd.Println(gcmd.Info, "接收到服务器消息:", string(msg))
	p := &packet.Package{}

	b, _ := p.Pack([]byte("123"))
	conn.Write(b)
}

func (this *client) OnSendMessage(conn net.Conn) {

}

// //客户端
func clientGo() {
	x := snetclient.NewClient("127.0.0.1", "496", &client{})
	x.Serve()
}
