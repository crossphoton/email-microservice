package main

import (
	"fmt"
	"log"
	"net"

	"github.com/crossphoton/email-microservice/src"
	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

// Config stores service configuration
type Config struct {
	Port int `mapstructure:"PORT"`
}

var config Config

func init() {
	viper.AddConfigPath(".")
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
}

func main() {
	// Registering service
	server := grpc.NewServer()
	emailServer := src.EmailServer{}
	src.RegisterEmailServiceServer(server, &emailServer)

	// Listening to port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("cannot listen on %s : %v", config.Port, err)
	}
	defer listener.Close()

	// Starting server
	log.Printf("starting server on %v", config.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("cannor start server: %v", err)
	}
}
