package chat

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_chat-server/internal/converter"
	"github.com/saver89/microservices_chat-server/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	op := "chatServer.SendMessage"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	err := i.chatService.SendMessage(ctx, converter.ToMessageInfo(req))
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
