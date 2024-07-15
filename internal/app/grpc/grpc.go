package grpc

import (
	"fmt"
	"log/slog"
	"net"

	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is the gRPC server
type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New creates a new gRPC server
func New(
	log *slog.Logger,
	port int,
) *Server {

	s := grpc.NewServer()
	reflection.Register(s)
	server := NewChatServer(log)
	desc.RegisterChatV1Server(s, server)

	grpcServer := Server{
		log:        log,
		port:       port,
		gRPCServer: s,
	}

	return &grpcServer
}

// Run starts the gRPC server
func (s *Server) Run() error {
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

// Stop stops the gRPC server
func (s *Server) Stop() {
	const op = "app.grpc.Stop"

	s.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", s.port))

	s.gRPCServer.GracefulStop()
}
