package snet

import (
	"github.com/golang/protobuf/proto"
	"snet/snet.v2/proto"
)

type Package struct {
	Err error
	Msg *pb.Msg
}

func (p *Package) Pack() []byte {
	b, err := proto.Marshal(p.Msg)
	if err != nil {
		p.Err = err
		return nil
	}
	return b
}

func (p *Package) UnPack(buf []byte) *Package {
	p.Err = proto.Unmarshal(buf, p.Msg)
	if p.Err != nil {
		return nil
	}
	return p
}
