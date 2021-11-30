package src

import (
	"context"
	"fmt"
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

var smtpClient *mail.SMTPClient

type EmailServer struct {
	UnimplementedEmailServiceServer
}

func (s *EmailServer) SendEmail(ctx context.Context, req *SendEmailRequest) (*ResponseMessage, error) {
	email := mail.NewMSG()

	// Add receivers
	email.
		AddTo(req.Recipients.To...).
		AddCc(req.Recipients.Cc...).
		AddBcc(req.Recipients.Bcc...).
		SetSubject(req.Subject)

	// Set content type
	if req.ContentType == "text/html" {
		email.SetBody(mail.TextHTML, req.GetBody())
	} else {
		email.SetBody(mail.TextPlain, req.GetBody())
	}

	// Check for errors
	if email.Error != nil {
		return nil, fmt.Errorf("email format error: %v", email.Error)
	}

	// Send the email
	err := email.Send(smtpClient)
	if err != nil {
		log.Printf("err: error sending email: %v", err)
		return nil, fmt.Errorf("error sending email: %v", err)
	}

	return &ResponseMessage{Success: true, Ack: "mail sent successfully"}, nil
}

func (s *EmailServer) SendRawEmail(ctx context.Context, req *RawSendEmailRequest) (*ResponseMessage, error) {
	// Sending email
	err := mail.SendMessage(config.SMTPEmail, req.Recipients, string(req.Body), smtpClient)
	if err != nil {
		log.Printf("err: error sending email: %v", err)
		return nil, fmt.Errorf("error sending email: %v", err)
	}

	return &ResponseMessage{Ack: "email sent successfully", Success: true}, nil
}

func (s *EmailServer) SendEmailWithTemplate(ctx context.Context, req *SendEmailWithTemplateRequest) (*ResponseMessage, error) {
	return nil, fmt.Errorf("not implemented")
}
