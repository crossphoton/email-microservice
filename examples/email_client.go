package main

import (
	"context"
	"log"
	"time"

	"github.com/crossphoton/email-microservice/src"
	"google.golang.org/grpc"
)

const address = "localhost:55055"

var client src.EmailServiceClient

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()

	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error connecting to server: %v", err)
	}
	defer connection.Close()
	client = src.NewEmailServiceClient(connection)
}

func main() {
	sendEmail("Hey there!!", []string{"crossphoton@gmail.com"})
}

func sendEmail(body string, receivers []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.SendEmail(ctx, &src.SendEmailRequest{Receipients: receivers, Body: body})
	if err != nil {
		log.Fatalf("request failed: %v", err)
	} else {
		log.Print(res)
	}
}
