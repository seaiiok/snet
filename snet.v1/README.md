# snet
# 这是一个简易的Tcp框架
 

## snet.v1 适用于十分简陋的任务

```Go

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

	for i := 0; i < 3; i++ {
		clientGo(byte(i))
	}
	time.Sleep(1 * time.Second)
}

func BenchmarkSnet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clientGo(0)
	}
}

func serverGo() {
	s := New()
	s.OnConnect(func(conn *Connection) {
		//建立连接事件
	})

	s.OnDisConnect(func(conn *Connection) {
		//断开连接事件
	})

	s.OnSendMessage(func(conn *Connection) {
		//向客户端发送数据
		// conn.OnSendMsg(snet.Package{})
	})

	s.OnRecvMessage(func(conn *Connection, msg Package) {
		//收到客户端数据
		//fmt.Println(msg)
		conn.OnSendMsg(msg)
	})
}

//建一个客户端
func clientGo(id byte) {
	conn, err := net.Dial("tcp", "127.0.0.1:495")
	if err != nil {
		fmt.Println("client dial err, exit!")
		return
	}

	//封包
	msg1 := makeSomeMsg(id)
	b := msg1.Pack()

	_, err = conn.Write(b)
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 1024)
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	//解包
	msg2 := Package{}
	msg := msg2.UnPack(buf[:cnt])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------接收数据------------")
	fmt.Println("ID:", msg.ID)
	fmt.Println("Key长度:", msg.KeyLen, "Key内容:", msg.Key)
	fmt.Println("Data长度:", msg.DataLen, "Data内容:", msg.Data)

}

//制造一些数据
func makeSomeMsg(id byte) Package {
	return Package{
		ID:      id,
		KeyLen:  0,
		DataLen: 0,
		Key:     []string{"golang", "tcp"},
		Data:    [][]string{{"1", "data1"}, {"2", "data2"}, {"2", "data2"}},
	}
}



```
result
  === RUN   TestSnet
  ------------接收数据------------
  ID: 0
  Key长度: 29 Key内容: [golang tcp]
  Data长度: 59 Data内容: [[1 data1] [2 data2] [2 data2]]
  ------------接收数据------------
  ID: 1
  Key长度: 29 Key内容: [golang tcp]
  Data长度: 59 Data内容: [[1 data1] [2 data2] [2 data2]]
  ------------接收数据------------
  ID: 2
  Key长度: 29 Key内容: [golang tcp]
  Data长度: 59 Data内容: [[1 data1] [2 data2] [2 data2]]
  --- PASS: TestSnet (1.00s)


