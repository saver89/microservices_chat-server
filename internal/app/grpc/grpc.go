package grpc

import (
	"fmt"
	"log/slog"
	"net"

	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *GRPCServer {

	s := grpc.NewServer()
	reflection.Register(s)
	server := NewChatServer(log)
	desc.RegisterChatV1Server(s, server)

	grpcServer := GRPCServer{
		log:        log,
		port:       port,
		gRPCServer: s,
	}

	return &grpcServer
}

func (s *GRPCServer) Run() error {
	const op = "app.grpc.Run"
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *GRPCServer) Stop() {
	const op = "app.grpc.Stop"

	s.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", s.port))

	s.gRPCServer.GracefulStop()
}
