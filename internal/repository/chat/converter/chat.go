package converter

import (
	"github.com/saver89/microservices_chat-server/internal/model"
	modelRepo "github.com/saver89/microservices_chat-server/internal/repository/chat/model"
)

// ToChatFromRepo converts a chat from repository to model
func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID: chat.ID,
		Info: model.ChatInfo{
			Name: chat.Name,
		},
		CreatedAt: chat.CreatedAt,
	}
}
