package snet

import (
	"io"
	"net"
)

func (s *Snet) NewConnection(conn *net.TCPConn) {
	defer conn.Close()
	go s.Conn.OnSendMessage(conn)
	for {
		msg := NextConnLength(conn)
		go s.Conn.OnRecvMessage(conn, msg)
	}

}

func NextConnLength(conn *net.TCPConn) (msg []byte) {
	p := &Package{}
	tempBuff := make([]byte, 0)

	for {
		msg = make([]byte, 0)
		oneByte := make([]byte, 1)
		n, err := conn.Read(oneByte)
		if err != nil || n != 1 {
			if err == io.EOF {
				//TODO client close
				return
			}
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
		if err != nil || n != int(tempMsgLength) {
			if err == io.EOF {
				//TODO client close
				return
			}
			continue
		}
		msg = append(msg, msgBytes...)
		return

	}
}
