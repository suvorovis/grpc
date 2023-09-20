package service2

import (
	"context"
	"log"
)

type Service2 struct {
	UnimplementedTestServer
}

func (s *Service2) Test(_ context.Context, req *SomeReq) (*SomeRes, error) {
	log.Printf("got request: %s", req.GetDesc())

	return &SomeRes{Desc: "yeah v1.2!"}, nil
}
