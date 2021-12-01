package src

import (
	"context"
	"fmt"
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

var smtpClient *mail.SMTPClient

// EmailServer implements EmailServiceServer
// to be used to create a gRPC server
type EmailServer struct {
	UnimplementedEmailServiceServer
}

// SendEmail sends an email with given Recipients, Subject, Body, .....
func (s *EmailServer) SendEmail(ctx context.Context, req *SendEmailRequest) (*ResponseMessage, error) {
	if err := req.Recipients.validate(); err != nil {
		return nil, err
	}

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

	// Add attachments
	for _, attachment := range req.GetAttachments() {
		email.AddAttachmentBase64(attachment.Base64Data, attachment.Filename)
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

// SendRawEmail sends an email with given Recipients and RFC822 formatted message
func (s *EmailServer) SendRawEmail(ctx context.Context, req *RawSendEmailRequest) (*ResponseMessage, error) {
	if err := validateEmails(req.Recipients); err != nil {
		return nil, err
	}

	// Sending email
	err := mail.SendMessage(config.SMTPEmail, req.Recipients, string(req.Body), smtpClient)
	if err != nil {
		log.Printf("err: error sending email: %v", err)
		return nil, fmt.Errorf("error sending email: %v", err)
	}

	return &ResponseMessage{Ack: "email sent successfully", Success: true}, nil
}

// SendEmailWithTemplate sends an email with email template
// and its parameters. Template should exists beforehand in ./template folder
func (s *EmailServer) SendEmailWithTemplate(ctx context.Context, req *SendEmailWithTemplateRequest) (*ResponseMessage, error) {
	// parsing template
	message, err := getMessageFromTemplate(req.GetTemplateName(), req.GetTemplateParams())
	if err != nil {
		return nil, err
	}

	// validating Recipients
	if err := req.Recipients.validate(); err != nil {
		return nil, err
	}

	// forming email
	email := mail.NewMSG()

	// Add receivers
	email.
		AddTo(req.Recipients.To...).
		AddCc(req.Recipients.Cc...).
		AddBcc(req.Recipients.Bcc...).
		SetSubject(req.Subject).
		SetBodyData(mail.TextHTML, message)

	// Add attachments
	for _, attachment := range req.GetAttachments() {
		email.AddAttachmentBase64(attachment.Base64Data, attachment.Filename)
	}

	// Check for errors
	if email.Error != nil {
		return nil, fmt.Errorf("email format error: %v", email.Error)
	}

	// Send the email
	err = email.Send(smtpClient)
	if err != nil {
		log.Printf("err: error sending email: %v", err)
		return nil, fmt.Errorf("error sending email: %v", err)
	}

	return &ResponseMessage{Success: true, Ack: "mail sent successfully"}, nil
}
