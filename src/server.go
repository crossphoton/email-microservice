package src

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
)

type EmailServer struct {
	UnimplementedEmailServiceServer
}

func (s *EmailServer) SendEmail(ctx context.Context, req *SendEmailRequest) (*ResponseMessage, error) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Sender data.
	from := "email@gmail.com"
	password := "password"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		auth,
		from,
		req.GetReceipients(),
		[]byte(req.GetBody()),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ResponseMessage{Success: true, Ack: "mail sent successfully"}, nil
}
