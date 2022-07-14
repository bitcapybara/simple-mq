package server

import (
	"context"

	"github.com/bitcapybara/geckod-proto/gen/go/proto/signaling"
	"github.com/bitcapybara/geckod/broker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	broker broker.Broker
	signaling.UnimplementedSignalingServer
}

func NewServer() *Server {
	return nil
}

func (s *Server) Run(ctx context.Context) {

}

func (s *Server) Register(grpcServer *grpc.Server) {
	signaling.RegisterSignalingServer(grpcServer, s)
}

func (s *Server) Ping(context.Context, *signaling.Empty) (*signaling.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func (s *Server) Signaling(signaling.Signaling_SignalingServer) error {
	return status.Errorf(codes.Unimplemented, "method Signaling not implemented")
}
