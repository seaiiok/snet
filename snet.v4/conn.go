package snet

import (
	"context"
	"io"
	"net"

	"snet/snet.v4/packet"
)

func (this *snet) newConnection(conn *net.TCPConn) {
	// defer conn.Close()
	childCtx, childCancel := context.WithCancel(this.ctx)
	go this.conn.OnSendMessage(conn)

	for {
		select {
		case <-childCtx.Done():
			defer conn.Close()
			return
		default:
		}

		msg := this.remoteConnHandle(childCtx, childCancel, conn)
		if len(msg) == 0 {
			continue
		}
		go this.conn.OnRecvMessage(conn, msg)

	}
}

func (this *snet) remoteConnHandle(ctx context.Context, cancel context.CancelFunc, conn *net.TCPConn) (msg []byte) {
	// defer conn.Close()

	Onebyte := make([]byte, 1)
	headBytes := make([]byte, 0)
	endByte := make([]byte, 1)

	p := &packet.Package{}

	for {

		select {
		case <-ctx.Done():
			return
		default:
		}

		n, err := conn.Read(Onebyte)
		if err != nil {
			if err == io.EOF {
				//TODO client close
				this.conn.OnDisConnect(conn, "client connection close")
				cancel()
				return []byte{}
			}
			this.conn.OnDisConnect(conn, err.Error())
			cancel()
			return []byte{}
		}

		if n != 1 {
			continue
		}

		headBytes = append(headBytes, Onebyte...)

		if len(headBytes) < 9 {
			continue
		}

		//验证起始符 -- '!'
		if headBytes[0] != 33 {
			headBytes = headBytes[1:9]
			continue
		}

		msgLength := p.UnPackMsgLength(headBytes[1:5])
		msgLengthCRC1 := p.CheckCRC32(msgLength)

		msgLengthCRC2 := p.UnPackMsgLength(headBytes[5:9])

		if msgLengthCRC1 != msgLengthCRC2 {
			headBytes = headBytes[1:9]
			continue
		}

		msg = make([]byte, msgLength)
		n, err = io.ReadFull(conn, msg)
		if err != nil {
			if err == io.EOF {
				//TODO client close
				this.conn.OnDisConnect(conn, "client connection close")
				cancel()
				return []byte{}
			}
			this.conn.OnDisConnect(conn, err.Error())
			cancel()
			return []byte{}
		}

		if uint32(n) != msgLength {
			return []byte{}
		}

		n, err = conn.Read(endByte)
		if err != nil {
			if err == io.EOF {
				//TODO client close
				this.conn.OnDisConnect(conn, "client connection close")
				cancel()
				return []byte{}
			}
			this.conn.OnDisConnect(conn, err.Error())
			cancel()
			return []byte{}
		}

		//验证结束符 -- '\n'
		if n != 1 || endByte[0] != 10 {
			return []byte{}
		}
		return
	}
}
