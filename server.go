// server.go
package handler

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/muzzapp/devops-interview-task/pkg/muzz"
	"github.com/muzzapp/devops-interview-task/server/internal"
)

type Server struct {
	muzz.UnimplementedServiceServer
}

func (s Server) Echo(ctx context.Context, req *muzz.EchoRequest) (*muzz.EchoResponse, error) {
	start := time.Now()
	defer func() {
		internal.echoRequestsTotal.Inc()
		internal.echoRequestDuration.Observe(time.Since(start).Seconds())
		internal.echoRequestDurationSummary.WithLabelValues("Echo").Observe(time.Since(start).Seconds())
	}()

	time.Sleep(time.Duration(rand.IntN(5000)) * time.Millisecond)

	// Simulate an error for demonstration purposes
	// Simulate a random error 10% of the time
	if rand.Intn(10) == 0 {
		internal.echoRequestsFailedTotal.Inc()
		internal.echoRequestsCounter.WithLabelValues("Echo", "error").Inc()
		return nil, fmt.Errorf("random error occurred")
	}
   
	// Increment the requests counter with the "success" status label
	internal.echoRequestsCounter.WithLabelValues("Echo", "success").Inc()
	return &muzz.EchoResponse{Message: req.GetMessage()}, nil
}