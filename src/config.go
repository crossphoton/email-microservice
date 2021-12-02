package src

import (
	"log"
	"os"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// Config stores SMTP configurations to use for sending emails
type Config struct {
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPEmail    string `mapstructure:"SMTP_EMAIL"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
}

var config Config
var err error

func init() {
	config.SMTPEmail = os.Getenv("SMTP_EMAIL")
	config.SMTPHost = os.Getenv("SMTP_HOST")
	config.SMTPPassword = os.Getenv("SMTP_PASSWORD")
	config.SMTPPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))

	smtpServer := mail.NewSMTPClient()
	smtpServer.Host = config.SMTPHost
	smtpServer.Port = config.SMTPPort
	smtpServer.Username = config.SMTPEmail
	smtpServer.Password = config.SMTPPassword
	smtpServer.Encryption = mail.EncryptionSTARTTLS

	smtpServer.KeepAlive = true
	smtpServer.ConnectTimeout = 10 * time.Second
	smtpServer.SendTimeout = 10 * time.Second
	// smtpServer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	smtpClient, err = smtpServer.Connect()
	if err != nil {
		log.Fatalf("connection to remote smtp server failed: %v", err)
	}
}
