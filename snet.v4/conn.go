package snet

import (
	"context"
	"io"
	"net"
)

func (s *snet) newConnection(ctx context.Context, conn *net.TCPConn) {
	// defer conn.Close()
	childCtx, childCancel := context.WithCancel(ctx)
	go s.conn.OnSendMessage(conn)

	remoteIP := conn.RemoteAddr().String()[:9]

	for {

		select {
		case <-childCtx.Done():
			defer conn.Close()
			return
		default:
		}
		switch remoteIP {
		case "-127.0.0.1":
			cmd, msg := s.localConnHandle(childCtx, childCancel, conn)
			if len(cmd) != 3 {
				continue
			}
			go s.conn.OnLocalCommand(conn, cmd, msg)

		default:
			msg := s.remoteConnHandle(childCtx, childCancel, conn)
			if len(msg) == 0 {
				continue
			}
			go s.conn.OnRecvMessage(conn, msg)
		}
	}
}

func (s *snet) remoteConnHandle(ctx context.Context, cancel context.CancelFunc, conn *net.TCPConn) (msg []byte) {
	// defer conn.Close()
	p := &Package{}
	tempBuff := make([]byte, 0)

	for {

		select {
		case <-ctx.Done():
			defer conn.Close()
			return
		default:
		}

		msg = make([]byte, 0)
		oneByte := make([]byte, 1)
		n, err := conn.Read(oneByte)
		if err != nil {
			if err == io.EOF {
				//TODO client close
				s.conn.OnDisConnect(conn, "client connection close")
				cancel()
				return
			}
			s.conn.OnDisConnect(conn, err.Error())
			cancel()
			return
		}

		if n != 1 {
			continue
		}

		tempBuff = append(tempBuff, oneByte...)
		tempBuffLength := len(tempBuff)
		if tempBuffLength < 8 {
			continue
		}

		tempBuff = tempBuff[tempBuffLength-8:]
		tempMsgLength := p.UnPackMsgLength(tempBuff[:4])
		tempMsgLengthCRC1 := p.CheckCRC32(tempMsgLength)
		tempMsgLengthCRC2 := p.UnPackMsgLength(tempBuff[4:8])
		if tempMsgLengthCRC1 != tempMsgLengthCRC2 {
			continue
		}

		msgBytes := make([]byte, tempMsgLength)
		n, err = io.ReadFull(conn, msgBytes)
		if err != nil {
			if err == io.EOF {
				//TODO client close
				s.conn.OnDisConnect(conn, "client connection close")
				cancel()
				return
			}
			s.conn.OnDisConnect(conn, err.Error())
			cancel()
			return
		}

		if n != int(tempMsgLength) {
			continue
		}

		msg = append(msg, msgBytes...)
		return

	}
}

func (s *snet) localConnHandle(ctx context.Context, cancel context.CancelFunc, conn *net.TCPConn) (cmd []byte, msg []byte) {
	// defer conn.Close()
	cmd = make([]byte, 0)
	msg = make([]byte, 0)
	cmdMsg := s.remoteConnHandle(ctx, cancel, conn)
	if len(cmdMsg) < 4 {
		return
	}
	if cmdMsg[3] != byte(32) {
		return
	}
	cmd = append(cmd, cmdMsg[:3]...)
	msg = append(msg, cmdMsg[4:]...)
	return
}
