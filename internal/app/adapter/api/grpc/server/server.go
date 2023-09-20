package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	service1v1 "github.com/suvorovis/grpc/internal/app/adapter/api/grpc/v1/service1"
	service2v1 "github.com/suvorovis/grpc/internal/app/adapter/api/grpc/v1/service2"
	service1v2 "github.com/suvorovis/grpc/internal/app/adapter/api/grpc/v2/service1"
	service2v2 "github.com/suvorovis/grpc/internal/app/adapter/api/grpc/v2/service2"
)

type Server struct {
	grpc *grpc.Server
}

func NewServer() *Server {
	s := &Server{}

	s.grpc = grpc.NewServer()
	service1v1.RegisterTestServer(s.grpc, &service1v1.Service1{})
	service2v1.RegisterTestServer(s.grpc, &service2v1.Service2{})
	service1v2.RegisterTestServer(s.grpc, &service1v2.Service1{})
	service2v2.RegisterTestServer(s.grpc, &service2v2.Service2{})

	return s
}

func (s *Server) Start(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("can't listen tcp port %d: %w", port, err)
	}

	log.Printf("server is listening on %v", listener.Addr())
	if err = s.grpc.Serve(listener); err != nil {
		return fmt.Errorf("can't run grpc server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		s.grpc.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
	case <-ctx.Done():
		log.Printf("can't stop gracefully")

		s.grpc.Stop()
	}
}
