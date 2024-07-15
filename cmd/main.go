package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/saver89/microservices_chat-server/internal/app"
)

const grpcPort = 50051

func main() {
	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	application := app.New(log, grpcPort)
	go func() {
		err := application.GRPCServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
}
