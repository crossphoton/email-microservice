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

func main() {
	sendEmail("Hey there!!", []string{"adiag1200@gmail.com", "cs19b1003@iiitr.ac.in"})
}

func sendEmail(body string, receivers []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.SendRawEmail(ctx, &src.RawSendEmailRequest{Recipients: receivers, Body: []byte(body)})
	if err != nil {
		log.Fatalf("request failed: %v", err)
	} else {
		log.Print(res)
	}
}
