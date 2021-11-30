package main

import (
	"log"
	"net"

	"github.com/crossphoton/email-microservice/src"
	grpc "google.golang.org/grpc"
)

const (
	PORT = ":55055"
)

func main() {
	// Registering service
	server := grpc.NewServer()
	emailServer := src.EmailServer{}
	src.RegisterEmailServiceServer(server, &emailServer)

	// Listening to port
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("cannot listen on %s : %v", PORT, err)
	}
	defer listener.Close()

	// Starting server
	log.Printf("starting server on %v", PORT)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("cannor start server: %v", err)
	}
}
