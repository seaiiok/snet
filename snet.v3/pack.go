package snet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gcom/gcmd"
	"hash/crc32"
)

//Package 消息包
type Package struct {
	Msg []byte
}

//Pack 封包
func (p *Package) Pack() []byte {

	buf := new(bytes.Buffer)

	msgLength := uint32(len(p.Msg))

	err := binary.Write(buf, binary.BigEndian, msgLength)
	if err != nil {
		gcmd.Println(gcmd.Err, "pack msg length err:", err)
		return nil
	}
	checkLength := p.CheckCRC32(msgLength)
	err = binary.Write(buf, binary.BigEndian, checkLength)
	if err != nil {
		gcmd.Println(gcmd.Err, "pack msg check length err:", err)
		return nil
	}

	err = binary.Write(buf, binary.BigEndian, p.Msg)
	if err != nil {
		gcmd.Println(gcmd.Err, "pack msg err:", err)
		return nil
	}

	return buf.Bytes()
}

// UnPack 解包
func (p *Package) UnPack(buff []byte) *Package {
	buf := bytes.NewBuffer(buff)

	var msgLength uint32
	err := binary.Read(buf, binary.BigEndian, &msgLength)
	if err != nil {
		gcmd.Println(gcmd.Err, "unpack msg length err:", err)
		return &Package{}
	}

	var checkLength uint32
	err = binary.Read(buf, binary.BigEndian, &checkLength)
	if err != nil {
		gcmd.Println(gcmd.Err, "unpack msg check length err:", err)
		return &Package{}
	}

	msg := make([]byte, msgLength)
	err = binary.Read(buf, binary.BigEndian, &p.Msg)
	if err != nil {
		gcmd.Println(gcmd.Err, "unpack msg err:", err)
		return &Package{}
	}

	p.Msg = append(p.Msg, msg...)

	return p
}

// UnPackMsgLength 解包长度
func (p Package) UnPackMsgLength(buff []byte) (length uint32) {
	buf := bytes.NewBuffer(buff)
	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		fmt.Println("unpack length err:", err)
		return 0
	}
	return
}

// SetMsg 封包消息
func (p *Package) SetMsg(msg []byte) {
	p.Msg = append(p.Msg, msg...)
}

// CheckCRC32 CRC校验
func (p *Package) CheckCRC32(l interface{}) uint32 {
	ieee := crc32.NewIEEE()
	binary.Write(ieee, binary.BigEndian, l)
	return ieee.Sum32()
}
