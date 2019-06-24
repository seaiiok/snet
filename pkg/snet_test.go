package pkg

import (
	"snet/snet/proto"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestConf(t *testing.T) {
	conf:=LoadConfigFile("./snet.json")
	t.Log(conf)
	
	u1 := &pb.Row{
		StdID:   "a",
		KeyNo:   "b",
		Status1: "c",
		Status2: "d",
	}
	u2 := &pb.Msg{}
	u2.Datas = append(u2.Datas, u1)
	b, err := proto.Marshal(u2)
	if err != nil {
		t.Log(err)
	}

	u3 := &pb.Msg{}

	err = proto.Unmarshal(b, u3)
	if err != nil {
		t.Log(err)
	}
	t.Log(u3.Datas)


}
