package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/homayoonalimohammadi/go-post-app/postapp/gen/go/postapp"
	"github.com/homayoonalimohammadi/go-post-app/postapp/internal/app/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	address := "0.0.0.0:8888"
	postappService := core.New(log)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggerUnaryInterceptor),
		grpc.StreamInterceptor(loggerStreamInterceptor),
	)
	postapp.RegisterPostAppServer(grpcServer, postappService)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	reflection.Register(grpcServer)

	log.Infof("Starting grpcServer on: %s", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func loggerUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("---> unary interceptor:", info.FullMethod)
	return handler(ctx, req)
}

func loggerStreamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("---> stream interceptor:", info.FullMethod)
	return handler(srv, stream)
}
