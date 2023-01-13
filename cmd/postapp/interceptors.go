package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type JWTInterceptor struct{}

func NewJWTInterceptor() *JWTInterceptor {
	return &JWTInterceptor{}
}

func (i *JWTInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("---> unary JWT interceptor")
		return handler(ctx, req)
	}
}

func (i *JWTInterceptor) Stream() grpc.StreamServerInterceptor {

	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("---> stream JWT interceptor")
		return handler(srv, stream)
	}
}
