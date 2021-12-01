package main

import (
	"context"
	"log"
	"time"

	"github.com/crossphoton/email-microservice/src"
)

func sendEmailStd(req *src.SendEmailRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.SendEmail(ctx, req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
		return err
	}

	log.Print(res)
	return nil
}
