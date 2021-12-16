package main

import (
	"context"
	"log"
	"time"

	"github.com/crossphoton/email-microservice/examples/src"
)

func sendEmailStd(req *src.SendEmailRequest) (*src.ResponseMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.SendEmail(ctx, req)
	if err != nil {
		log.Printf("request failed: %v", err)
		return nil, err
	}

	return res, nil
}

func StdEmail() (*src.ResponseMessage, error) {
	emailRequest := src.SendEmailRequest{
		Recipients: &src.Recipients{
			To: []string{
				"Aditya Agrawal<email@e.crossphoton.tech>",
			},
			Cc: []string{
				"Spam Address<spam@e.crossphoton.tech>",
			},
			Bcc: []string{
				"support@e.crossphoton.tech",
			},
		},
		Subject:     "Hi there. I hope you're good",
		ContentType: "text/html",
		Body:        "<h1>This is heading</h1><p>This is text</p>",
	}

	return sendEmailStd(&emailRequest)
}
