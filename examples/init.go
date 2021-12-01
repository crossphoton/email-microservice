package main

import (
	"context"
	"log"
	"time"

	"github.com/crossphoton/email-microservice/src"
	"google.golang.org/grpc"
)

const address = "localhost:5555"

var client src.EmailServiceClient

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()

	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error connecting to server: %v", err)
	}
	client = src.NewEmailServiceClient(connection)
}
