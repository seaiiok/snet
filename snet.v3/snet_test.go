package snet

import (
	"fmt"
	"gcom/gcmd"
	"gcom/gfiles"
	"net"
	"testing"
	"time"

	"snet/snet.v3/iface"
)

func TestSnet(t *testing.T) {

	//启动Tcp服务器
	go serverGo()

	time.Sleep(1 * time.Second)

	//启动客户端
	// go clientGo()

	select {}
}

// func BenchmarkSnet(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		clientGo()
// 	}
// }

type Client struct {
	svr iface.ISnet
}

func serverGo() {
	s := NewServer("localhost", "496")
	c := &Client{svr: s}

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
	// gcmd.Println(gcmd.Ok, "接收客户端消息:", msg)
	gcmd.Println(gcmd.Ok, "接收客户端消息:", string(msg))
}

func (c *Client) OnSendMessage(conn *net.TCPConn) {

}

func (c *Client) OnLocalCommand(conn *net.TCPConn, cmd []byte, msg []byte) {
	gcmd.Println(gcmd.Ok, "本地命令:", string(cmd), string(msg))
}

// //客户端
func clientGo() {
	conn, err := net.Dial("tcp", ":496")
	if err != nil {
		gcmd.Println(gcmd.Warn, "client dial err, exit!")
		return
	}

	time.Sleep(1 * time.Second)
	gcmd.Println(gcmd.Ok, "客户端", conn.LocalAddr().String(), conn.RemoteAddr().String())
	p := &Package{}
	// b := make([]byte, 0)
	//本机测试超过1.5G的数据通讯加写入本地文件总用时超过20秒,异地机的话没测试....肯定无语
	// for i := 0; i < 20000000; i++ {
	// 	b = append(b, []byte("123456789 的艰苦拉萨机法师打房价开始打附件卡萨丁就")...)
	// }
	// b = append(b, []byte("###")...)
	// fmt.Println("发送字节长度:", len(b))

	b1, err := gfiles.GetFileMD5("./KNR04886_012_1.zip")
	gcmd.Println(gcmd.Warn, fmt.Sprintf("%x", b1), err)

	// b, err := ioutil.ReadFile("./KNR04886_012_1.zip")
	// r, _ := gfiles.("./KNR04886_012_1.zip")

	// gcmd.Println(gcmd.Warn, b, err)
	// p.Msg = r
	x := p.Pack()

	conn.Write(x)

	defer conn.Close()

}

/*
=== RUN   TestSnet
客户端连接: 127.0.0.1:51864
客户端 127.0.0.1:51864 127.0.0.1:496
发送字节长度: 1400000003
接收客户端消息长度: 1400000003
客户端断开连接: 127.0.0.1:51864 原因: client connection close
接收客户端消息: ###
写入文件: <nil> 1400000003
exit status 2
FAIL    snet/snet.v3    21.330s
*/
