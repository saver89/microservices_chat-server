package converter

import (
	"github.com/saver89/microservices_chat-server/internal/model"
	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
)

// ToChatInfo converts a chat info from request to model
func ToChatInfo(req *desc.CreateRequest) *model.ChatInfo {
	return &model.ChatInfo{
		Name:      req.Name,
		UserNames: req.Usernames,
	}
}

// ToMessageInfo converts a message info from request to model
func ToMessageInfo(req *desc.SendMessageRequest) *model.MessageInfo {
	return &model.MessageInfo{
		FromUser: req.From,
		ChatID:   req.ChatId,
		Text:     req.Text,
		SentAt:   req.Timestamp.AsTime(),
	}
}
