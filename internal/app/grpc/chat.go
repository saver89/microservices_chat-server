package grpc

import (
	"context"
	"log/slog"

	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type chatServer struct {
	log *slog.Logger
	desc.UnimplementedChatV1Server
}

func NewChatServer(log *slog.Logger) *chatServer {
	return &chatServer{
		log: log,
	}
}

func (cs *chatServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	op := "chatServer.Create"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &desc.CreateResponse{}, nil
}

func (cs *chatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	op := "chatServer.SendMessage"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &emptypb.Empty{}, nil
}

func (cs *chatServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	op := "chatServer.Delete"
	cs.log.InfoContext(ctx, op, slog.Any("req", req))
	return &emptypb.Empty{}, nil
}
