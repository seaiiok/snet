package snet

import (
	"context"
	"fmt"
	"gcom/gcmd"
	"io"
	"net"

	"snet/snet.v4/iface"
)

type snet struct {
	ip     string
	port   int
	ctx    context.Context
	cancel context.CancelFunc
	conn   iface.ISnet
}

func NewServer(ip string, port int, s iface.ISnet) *snet {
	ctx, cancel := context.WithCancel(context.Background())
	return &snet{
		ip:     ip,
		port:   port,
		ctx:    ctx,
		cancel: cancel,
		conn:   s,
	}
}

func (s *snet) Serve() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		gcmd.Println(gcmd.Err, "server tcp addr err:", err)
		return
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		gcmd.Println(gcmd.Err, "server listen tcp err:", err)
		return
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())

	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done():
				gcmd.Println(gcmd.Err, "snet server exit:")
				defer l.Close()
				return
			default:
			}

			conn, err := l.AcceptTCP()
			if err != nil {
				gcmd.Println(gcmd.Err, "server accept tcp err:", err)
				// TODO 客户端断开连接
				if err == io.EOF {
					s.conn.OnDisConnect(conn, "client close connection")
					continue
				}
				s.conn.OnDisConnect(conn, err.Error())
				continue
			}
			// TODO 客户端连接
			s.conn.OnConnect(conn)

			// TODO 连接处理协程
			go s.newConnection(ctx, conn)

		}
	}(s.ctx)
}

func (s *snet) Stop() {
	s.cancel()
}
