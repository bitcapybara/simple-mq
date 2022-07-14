package main

import (
	"log"
	"net"

	"github.com/bitcapybara/simple-mq/config"
	"github.com/bitcapybara/simple-mq/server"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.FromYaml("./config/application.yaml")
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	server := server.NewServer()
	server.Register(grpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
