package main

import (
	"context"

	proto "answersys/ans.srv.v1.idgen/proto"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake
var set sonyflake.Settings

//:init
func Init() {
	sf = sonyflake.NewSonyflake(set)
	if sf == nil {
		panic("init sonyflake err")
	}
}

type IdGen struct{}

func (i *IdGen) GetId(ctx context.Context, in *proto.Req, out *proto.Response) error {
	set.MachineID = func() (uint16, error) {
		return uint16(in.MId), nil
	}
	if sf == nil {
		panic("sonyflake nil")
	}
	id, err := sf.NextID()
	if err != nil {
		return err
	}
	out.Id = id
	return nil
}
