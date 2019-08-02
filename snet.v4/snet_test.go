package snet

import (
	"bytes"
	"fmt"
	"gcom/gcmd"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestSnet(t *testing.T) {
	go serverGo()
	go clientGo()
	select {}
}

func serverGo() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "496"))
	if err != nil {
		gcmd.Println(gcmd.Err, "server tcp addr err:", err)
		return
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		gcmd.Println(gcmd.Err, "server listen tcp err:", err)
		return
	}

	go func() {
		conn, err := l.AcceptTCP()
		if err != nil {
			gcmd.Println(gcmd.Err, "server listen tcp err:", err)
			return
		}
		go connhandle(conn)
	}()

}

func clientGo() {
	conn, err := net.Dial("tcp", ":496")
	if err != nil {
		gcmd.Println(gcmd.Warn, "client dial err, exit!")
		return
	}
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		msg := strconv.Itoa(i)
		conn.Write([]byte(msg))
	}

}

func connhandle(conn *net.TCPConn) {
	bs := bytes.Buffer{}

	go func() {
		for {
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err != nil {
				gcmd.Println(gcmd.Warn, "client dial err, exit!")
			}

			bn, err := bs.Write(buf[:n])
			if err != nil || n != bn {
				gcmd.Println(gcmd.Warn, "client dial err, exit!")
			}
		}
	}()

	go func(){

	}()

}


