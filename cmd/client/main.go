package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/suvorovis/grpc/internal/app/adapter/api/grpc/v1/service1"
)

func main() {
	connection, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't dial: %s", err)
	}
	defer connection.Close()

	client := service1.NewTestClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Test(ctx, &service1.SomeReq{Desc: "asdfasdf"})
	if err != nil {
		log.Fatalf("can't send message: %s", err)
	}

	log.Printf("got response: %s", res.GetDesc())
}
