package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	healthpb "github.com/crossphoton/email-microservice/health"
	"github.com/crossphoton/email-microservice/src"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	grpc "google.golang.org/grpc"
)

// Config stores service configuration
type Config struct {
	Port           int `mapstructure:"PORT"`
	PrometheusPort int `mapstructure:"PROMETHEUS_PORT"`
}

var config Config

func init() {
	config.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	config.PrometheusPort, _ = strconv.Atoi(os.Getenv("PROMETHEUS_PORT"))
}

func main() {
	reg := prometheus.NewRegistry()
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf(":%d", config.PrometheusPort),
	}

	// Registering service
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)

	// Initialize all metrics.
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	grpcMetrics.InitializeMetrics(server)

	// Start your http server for prometheus.
	go func() {
		log.Println("starting server for prometheus at ", config.PrometheusPort)

		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("unable to start a http server for prometheus")
		}
	}()

	emailServer := src.EmailServer{}
	src.RegisterEmailServiceServer(server, &emailServer)
	healthpb.RegisterHealthServer(server, &emailServer)

	// Listening to port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("cannot listen on %d : %v", config.Port, err)
	}
	defer listener.Close()

	// Starting server
	log.Printf("starting server on %v", config.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("cannor start server: %v", err)
	}
}
