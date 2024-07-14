package app

import (
	"log/slog"

	"github.com/saver89/microservices_chat-server/internal/app/grpc"
)

type App struct {
	GRPCServer *grpc.GRPCServer
}

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
