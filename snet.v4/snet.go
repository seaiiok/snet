package snet // import "snet.v4"

import (
	"context"
	"fmt"
	"io"
	"net"
)

type snet struct {
	ip     string
	port   string
	ctx    context.Context
	cancel context.CancelFunc
	conn   ISnet
}

type ISnet interface {
	OnConnect(*net.TCPConn)
	OnDisConnect(*net.TCPConn, string)
	OnRecvMessage(*net.TCPConn, []byte)
	OnSendMessage(*net.TCPConn)
}

func NewServer(ip string, port string, s ISnet) *snet {
	ctx, cancel := context.WithCancel(context.Background())
	return &snet{
		ip:     ip,
		port:   port,
		ctx:    ctx,
		cancel: cancel,
		conn:   s,
	}
}

func (this *snet) Serve() error {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", this.ip, this.port))
	if err != nil {
		return err
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	go func() {

		for {
			select {
			case <-this.ctx.Done():
				defer l.Close()
			default:

			}

			conn, err := l.AcceptTCP()
			if err != nil {
				// TODO 客户端断开连接
				if err == io.EOF {
					this.conn.OnDisConnect(conn, "client close connection")
					continue
				}
				this.conn.OnDisConnect(conn, err.Error())
				continue
			}
			// TODO 客户端连接
			this.conn.OnConnect(conn)

			// TODO 连接处理协程
			go this.newConnection(conn)

		}
	}()
	return nil
}

func (this *snet) Stop() {
	this.cancel()
}
