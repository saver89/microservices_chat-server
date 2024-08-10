package chat

import (
	"github.com/saver89/microservices_chat-server/internal/client/db"
	"github.com/saver89/microservices_chat-server/internal/repository"
	"github.com/saver89/microservices_chat-server/internal/service"
)

const (
	createLog = "create"
	updateLog = "update"
	deleteLog = "delete"
	getLog    = "delete"
)

type serv struct {
	chatRepository     repository.ChatRepository
	chatUserRepository repository.ChatUserRepository
	chatLogRepository  repository.ChatLogRepository
	txManager          db.TxManager
}

// NewChatService creates a new chat service
func NewChatService(
	chatRepository repository.ChatRepository,
	chatUserRepository repository.ChatUserRepository,
	chatLogRepository repository.ChatLogRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository:     chatRepository,
		chatUserRepository: chatUserRepository,
		chatLogRepository:  chatLogRepository,
		txManager:          txManager,
	}
}
