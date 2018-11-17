package main

import (
	proto "answersys/ans.srv.v1.idgen/proto"
	"github.com/micro/go-micro"
)

func main() {
	Init()
	service := micro.NewService(
		micro.Name("id.gen"),
	)
	proto.RegisterIdGenServiceHandler(service.Server(), new(IdGen))
	service.Init()
	service.Run()
}
