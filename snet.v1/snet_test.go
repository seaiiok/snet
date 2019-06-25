package snet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	s := New("127.0.0.1:494")
	s.OnConnect(func(conn *net.TCPConn) {
		fmt.Println("连接：", conn.RemoteAddr().String())
	})

	s.OnDisConnect(func(conn *net.TCPConn) {
		fmt.Println("断开：", conn.RemoteAddr().String())
	})

	s.OnMessage(func(conn *net.TCPConn, msg string) {
		fmt.Println("消息：", msg)
	})

	for {
		time.Sleep(10 * time.Second)
	}
}
