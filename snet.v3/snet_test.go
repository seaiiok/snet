package snet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSnet(t *testing.T) {

	//启动Tcp服务器
	serverGo()

	//启动客户端
	clientGo()
	time.Sleep(1 * time.Second)
}

func BenchmarkSnet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clientGo()
	}
}

func serverGo() {
	s := New("localhost", "496")
	s.OnConnect(func(conn *Connection) {
		//建立连接事件
		fmt.Println("客服端建立连接:", conn.Conn.RemoteAddr())
	})

	s.OnDisConnect(func(conn *Connection) {
		//断开连接事件
		fmt.Println("客服端断开连接:", conn.Conn.RemoteAddr())
	})

	s.OnSendMessage(func(conn *Connection) {
		//向客户端发送数据
		// conn.OnSendMsg([]byte("msg..."))
	})

	s.OnRecvMessage(func(conn *Connection, msg []byte) {
		//收到客户端数据
		fmt.Println("服务器接收数据:", string(msg))
		msg = append(msg, []byte("Ok!")...)
		conn.OnSendMsg(msg)
	})
}

//客户端
func clientGo() {
	conn, err := net.Dial("tcp", "127.0.0.1:496")
	if err != nil {
		fmt.Println("client dial err, exit!")
		return
	}

	//制造Demo数据
	msg1 := makeSomeMsg()

	//封包
	pack1 := Package{}
	b1 := pack1.Pack(msg1)

	//发送数据
	_, err = conn.Write(b1)
	if err != nil {
		fmt.Println(err)
	}

	//接受数据
	buf := make([]byte, 1024)
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	//解包
	msg2 := Package{}
	b2 := msg2.UnPack(buf[:cnt])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("客服端接收数据:", string(b2))
}

//制造一些数据
func makeSomeMsg() []byte {
	return []byte("this is a test message!")
}
