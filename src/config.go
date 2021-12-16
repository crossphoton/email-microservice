package src

import (
	"flag"
	"os"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
	"go.uber.org/zap"
)

// Config stores SMTP configurations to use for sending emails
type Config struct {
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPEmail    string `mapstructure:"SMTP_EMAIL"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	disableEmail bool   `mapstructure:"DISABLE_EMAIL"`
}

var config Config
var err error
var logger *zap.Logger

func Initialize(_logger *zap.Logger) {
	logger = _logger

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

	var err error
	if !config.disableEmail {
		smtpClient, err = smtpServer.Connect()
	}
	if err != nil {
		logger.Fatal("failed to connect to SMTP server", zap.Error(err))
	}
}

func init() {
	// From environment variables
	config.SMTPEmail = os.Getenv("SMTP_EMAIL")
	config.SMTPHost = os.Getenv("SMTP_HOST")
	config.SMTPPassword = os.Getenv("SMTP_PASSWORD")
	config.SMTPPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	config.disableEmail, _ = strconv.ParseBool(os.Getenv("DISABLE_EMAIL"))

	// From flags
	flag.BoolVar(&config.disableEmail, "disableEmail", config.disableEmail, "disable email sending")
	flag.StringVar(&config.SMTPEmail, "smtpEmail", config.SMTPEmail, "SMTP email")
	flag.StringVar(&config.SMTPHost, "smtpHost", config.SMTPHost, "SMTP host")
	flag.StringVar(&config.SMTPPassword, "smtpPassword", config.SMTPPassword, "SMTP password")
	flag.IntVar(&config.SMTPPort, "smtpPort", config.SMTPPort, "SMTP port")
}
