package chat

import (
	"context"
	"log/slog"

	"github.com/saver89/microservices_chat-server/internal/log"
	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a chat
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	op := "chatServer.Delete"
	i.log.InfoContext(ctx, op, slog.Any("req", req))

	err := i.chatService.Delete(ctx, req.Id)
	if err != nil {
		i.log.ErrorContext(ctx, op, log.Err(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
