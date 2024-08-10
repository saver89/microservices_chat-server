package service

import (
	"context"

	"github.com/saver89/microservices_chat-server/internal/model"
)

// ChatService is the interface for chat service
type ChatService interface {
	Create(ctx context.Context, req *model.ChatInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Chat, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, req *model.MessageInfo) error
}

// ChatLogService is the interface for chat log service
type ChatLogService interface {
	Create(ctx context.Context, req model.ChatLogInfo) (int64, error)
}
