package packet

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

type Package struct {
}

//Pack 封包
func (this *Package) Pack(msg []byte) ([]byte, error) {

	buf := new(bytes.Buffer)

	startByte := []byte{33}
	err := binary.Write(buf, binary.BigEndian, startByte)
	if err != nil {
		return nil, err
	}

	msgLength := uint32(len(msg))
	err = binary.Write(buf, binary.BigEndian, msgLength)
	if err != nil {
		return nil, err
	}

	checkLength := this.CheckCRC32(msgLength)
	err = binary.Write(buf, binary.BigEndian, checkLength)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, msg)
	if err != nil {
		return nil, err
	}

	endByte := []byte{10}
	err = binary.Write(buf, binary.BigEndian, endByte)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnPack 解包
func (this *Package) UnPack(buff []byte) ([]byte, error) {
	buf := bytes.NewBuffer(buff)

	startByte := make([]byte, 1)
	err := binary.Read(buf, binary.BigEndian, &startByte)
	if err != nil {
		return nil, err
	}

	var msgLength uint32
	err = binary.Read(buf, binary.BigEndian, &msgLength)
	if err != nil {
		return nil, err
	}

	var checkLength uint32
	err = binary.Read(buf, binary.BigEndian, &checkLength)
	if err != nil {
		return nil, err
	}

	msg := make([]byte, msgLength)
	err = binary.Read(buf, binary.BigEndian, &msg)
	if err != nil {
		return nil, err
	}

	endByte := make([]byte, 1)
	err = binary.Read(buf, binary.BigEndian, &endByte)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// unPackMsgLength 解包长度
func (this *Package) UnPackMsgLength(buff []byte) (length uint32) {
	buf := bytes.NewBuffer(buff)
	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		return 0
	}
	return
}

// checkCRC32 CRC校验
func (this *Package) CheckCRC32(l interface{}) uint32 {
	ieee := crc32.NewIEEE()
	binary.Write(ieee, binary.BigEndian, l)
	return ieee.Sum32()
}
