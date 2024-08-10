package repository

import (
	"context"

	"github.com/saver89/microservices_chat-server/internal/model"
)

// ChatRepository is the interface for chat repository
type ChatRepository interface {
	Create(ctx context.Context, name string) (id int64, err error)
	Delete(ctx context.Context, id int64) (err error)
	Get(ctx context.Context, id int64) (chat *model.Chat, err error)
}

// ChatRepository is the interface for chat repository
type ChatUserRepository interface {
	Create(ctx context.Context, chatID int64, userNames []string) error
	Get(ctx context.Context, chatID int64) ([]string, error)
	Delete(ctx context.Context, chatID int64) error
}

// ChatLogRepository is the interface for chat log repository
type ChatLogRepository interface {
	Create(ctx context.Context, req model.ChatLogInfo) (int64, error)
}
