// main.go
package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/muzzapp/devops-interview-task/pkg/muzz"
	"github.com/muzzapp/devops-interview-task/server/internal"
	"github.com/muzzapp/devops-interview-task/server/internal/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()
	muzz.RegisterServiceServer(server, handler.Server{})

	// Run the gRPC server
	go func() {
		listener, listenerErr := net.Listen("tcp", "0.0.0.0:9876")
		if listenerErr != nil {
			slog.Error("Failed to create gRPC listener", "err", listenerErr)
			os.Exit(1)
		}

		slog.Info("gRPC server starting")
		if err := server.Serve(listener); err != nil {
			slog.Error("gRPC server stopped", "err", err)
			os.Exit(1)
		}
	}()

	// Expose metrics over HTTP
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		slog.Info("Metrics server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			slog.Error("Metrics server stopped", "err", err)
			os.Exit(1)
		}
	}()

	// Gracefully shut down gRPC server after receiving an interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	slog.Info("Shutting down gRPC server")
	server.GracefulStop()
}