package main

import (
	"log"

	server2 "github.com/suvorovis/grpc/internal/app/adapter/api/grpc/server"
)

func main() {
	server := server2.NewServer()
	if err := server.Start(9090); err != nil {
		log.Fatal(err)
	}
}
