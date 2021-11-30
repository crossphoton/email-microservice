package src

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/spf13/viper"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Config struct {
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPEmail    string `mapstructure:"SMTP_EMAIL"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
}

var (
	config Config
)

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			log.Fatalf("error reading config file: %v", err)
		}
	}

	viper.Unmarshal(&config)

	smtpServer := mail.NewSMTPClient()
	smtpServer.Host = config.SMTPHost
	smtpServer.Port = config.SMTPPort
	smtpServer.Username = config.SMTPEmail
	smtpServer.Password = config.SMTPPassword
	smtpServer.Encryption = mail.EncryptionSTARTTLS

	smtpServer.KeepAlive = true
	smtpServer.ConnectTimeout = 10 * time.Second
	smtpServer.SendTimeout = 10 * time.Second
	smtpServer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	smtpClient, err = smtpServer.Connect()
	if err != nil {
		log.Fatalf("connection to remote smtp server failed: %v", err)
	}
}
