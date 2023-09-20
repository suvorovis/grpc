package service1

import (
	"context"
	"log"
)

type Service1 struct {
	UnimplementedTestServer
}

func (s *Service1) Test(_ context.Context, req *SomeReq) (*SomeRes, error) {
	log.Printf("got request: %s", req.GetDesc())

	return &SomeRes{Desc: "yeah v1.1!"}, nil
}
