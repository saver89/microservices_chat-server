package chat

import (
	"log/slog"

	"github.com/saver89/microservices_chat-server/internal/service"
	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
)

// Implementation is the chat implementation
type Implementation struct {
	desc.UnimplementedChatV1Server

	log         *slog.Logger
	chatService service.ChatService
}

// NewImplementation creates a new chat implementation
func NewImplementation(log *slog.Logger, chatService service.ChatService) *Implementation {
	return &Implementation{
		log:         log,
		chatService: chatService,
	}
}
