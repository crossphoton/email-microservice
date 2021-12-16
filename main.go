package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	healthpb "github.com/crossphoton/email-microservice/health"
	"github.com/crossphoton/email-microservice/src"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	// grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

// Config stores service configuration
type Config struct {
	Port           int    `mapstructure:"PORT"`
	PrometheusPort int    `mapstructure:"PROMETHEUS_PORT"`
	Environment    string `mapstructure:"ENVIRONMENT"`
}

var config Config
var logger *zap.Logger

func init() {
	// From environment variables
	config.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	config.PrometheusPort, _ = strconv.Atoi(os.Getenv("PROMETHEUS_PORT"))

	// From command line
	flag.IntVar(&config.Port, "port", 5555, "port to listen")
	flag.IntVar(&config.PrometheusPort, "prometheusPort", 9090, "port to listen for prometheus")
	flag.StringVar(&config.Environment, "environment", "development", "environment")
	help := flag.Bool("help", false, "show help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	reg := prometheus.NewRegistry()
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf(":%d", config.PrometheusPort),
	}

	// Create logger based on environment
	var err error
	if config.Environment == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment(
			zap.AddStacktrace(zap.ErrorLevel),
			zap.AddStacktrace(zap.InfoLevel),
			zap.AddStacktrace(zap.FatalLevel),
			zap.AddStacktrace(zap.DebugLevel),
			zap.AddStacktrace(zap.DPanicLevel),
		)
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Registering service with middlewares
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_validator.UnaryServerInterceptor(),
				// grpc_opentracing.UnaryServerInterceptor(),
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(logger),
				grpc_validator.StreamServerInterceptor(),
				// grpc_opentracing.StreamServerInterceptor(),
			),
		),
	)

	// Initialize mail service
	src.Initialize(logger)

	// Initialize all metrics.
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	grpcMetrics.InitializeMetrics(server)

	// Start your http server for prometheus.
	go func() {
		logger.Info("starting server for prometheus", zap.Int("port", config.PrometheusPort))

		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error("failed to start http server for prometheus", zap.Error(err))
		}
	}()

	emailServer := src.EmailServer{}
	src.RegisterEmailServiceServer(server, &emailServer)
	healthpb.RegisterHealthServer(server, &emailServer)

	// Listening to port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.Fatal("cannot listen on port", zap.Int("port", config.Port), zap.Error(err))
	}
	defer listener.Close()

	// Starting server
	logger.Info("starting service", zap.Int("port", config.Port))
	if err := server.Serve(listener); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
