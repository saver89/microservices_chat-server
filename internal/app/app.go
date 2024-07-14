package app

import (
	"log/slog"

	"github.com/saver89/microservices_chat-server/internal/app/grpc"
)

// App is the main application
type App struct {
	GRPCServer *grpc.GRPCServer
}

// New creates a new application
func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	gRPCServer := grpc.New(log, grpcPort)

	app := App{
		GRPCServer: gRPCServer,
	}
	return &app
}
