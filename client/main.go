// main.go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"


	"github.com/muzzapp/devops-interview-task/pkg/muzz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, dialErr := grpc.Dial("127.0.0.1:9876", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if dialErr != nil {
		slog.Error("Failed to dial gRPC service: %v", "err", dialErr)
		os.Exit(1)
	}
	defer conn.Close()

	client := muzz.NewServiceClient(conn)

	// Send multiple Echo requests concurrently
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			response, respErr := client.Echo(
				context.Background(),
				&muzz.EchoRequest{Message: fmt.Sprintf("Hello, world! (%d)", i)},
			)
			if respErr != nil {
				slog.Error("Failed to call gRPC service: %v", "err", respErr)
				return
			}

			slog.Info("Response: %s", "message", response.Message)
		}()
	}

	wg.Wait()
	slog.Info("All requests completed")
}