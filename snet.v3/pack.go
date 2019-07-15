package snet

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//Package 消息包
type Package struct {
	DataLen int32
	Data    []byte
}

//Pack 封包
func (p Package) Pack(data []byte) []byte {
	buf := new(bytes.Buffer)

	p.DataLen = int32(len(data))
	p.Data = data

	err := binary.Write(buf, binary.LittleEndian, p.DataLen)
	if err != nil {
		fmt.Println("pack datalen:", err)
		return nil
	}

	err = binary.Write(buf, binary.LittleEndian, p.Data)
	if err != nil {
		fmt.Println("unpack data:", err)
		return nil
	}

	return buf.Bytes()
}

//UnPack 解包
func (p Package) UnPack(buf []byte) []byte {
	buff := bytes.NewBuffer(buf)

	err := binary.Read(buff, binary.LittleEndian, &p.DataLen)
	if err != nil {
		fmt.Println("unpack datalen:", err)
	}
	dataBuf := make([]byte, p.DataLen)
	err = binary.Read(buff, binary.LittleEndian, &dataBuf)
	if err != nil {
		fmt.Println("unpack data:", err)
	}

	return dataBuf
}
