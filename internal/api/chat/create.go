package chat

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_chat-server/internal/converter"
	"github.com/saver89/microservices_chat-server/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
)

// Create creates a chat
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	op := "chatServer.Create"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	id, err := i.chatService.Create(ctx, converter.ToChatInfo(req))
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}
	res := &desc.CreateResponse{
		Id: id,
	}

	return res, nil
}
