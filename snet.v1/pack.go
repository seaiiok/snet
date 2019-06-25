package snet

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
)

type Package struct {
	ID      byte
	KeyLen  int32
	DataLen int32
	Key     []string
	Data    [][]string
}

func (p Package) Pack() []byte {
	buf := new(bytes.Buffer)

	keyBuf := gobEn(p.Key)
	p.KeyLen = int32(len(keyBuf))

	dataBuf := gobEn(p.Data)
	p.DataLen = int32(len(dataBuf))

	err := binary.Write(buf, binary.LittleEndian, p.ID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = binary.Write(buf, binary.LittleEndian, p.KeyLen)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = binary.Write(buf, binary.LittleEndian, p.DataLen)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = binary.Write(buf, binary.LittleEndian, keyBuf)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = binary.Write(buf, binary.LittleEndian, dataBuf)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buf.Bytes()
}

func (p Package) UnPack(buf []byte) Package {
	buff := bytes.NewBuffer(buf)

	err := binary.Read(buff, binary.LittleEndian, &p.ID)
	if err != nil {
		fmt.Println(err)
	}

	err = binary.Read(buff, binary.LittleEndian, &p.KeyLen)
	if err != nil {
		fmt.Println(err)
	}

	err = binary.Read(buff, binary.LittleEndian, &p.DataLen)
	if err != nil {
		fmt.Println(err)
	}

	keyBuf := make([]byte, p.KeyLen)
	err = binary.Read(buff, binary.LittleEndian, &keyBuf)
	if err != nil {
		fmt.Println(err)
	}
	pKey := make([]string, 0)
	gobDe(keyBuf, &pKey)
	p.Key = pKey

	dataBuf := make([]byte, p.DataLen)
	err = binary.Read(buff, binary.LittleEndian, &dataBuf)
	if err != nil {
		fmt.Println(err)
	}
	pData := make([][]string, 0)
	gobDe(dataBuf, &pData)
	p.Data = pData

	return p
}

func gobEn(m interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func gobDe(m []byte, v interface{}) error {

	buf := bytes.NewBuffer(m)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(v)
	return err
}
