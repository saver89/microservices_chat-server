package converter

import (
	"github.com/saver89/microservices_chat-server/internal/model"
	desc "github.com/saver89/microservices_proto/pkg/chat/v1"
)

func ToChatInfo(req *desc.CreateRequest) *model.ChatInfo {
	return &model.ChatInfo{
		Name:      req.Name,
		UserNames: req.Usernames,
	}
}
