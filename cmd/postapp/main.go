package main

import (
	"io/ioutil"
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
	grpcServer := grpc.NewServer()
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
