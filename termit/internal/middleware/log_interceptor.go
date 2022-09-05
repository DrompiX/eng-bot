package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type logInterceptor struct{}

func NewLogInterceptor() *logInterceptor {
	return &logInterceptor{}
}

func (i *logInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		elapsed := time.Since(start)
		log.Printf("[%v, Elapsed: %s] | Request: %v ", info.FullMethod, elapsed, req)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		return resp, err
	}
}
