package chat

import (
	"context"

	"github.com/saver89/microservices_chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, req *model.ChatInfo) (int64, error) {
	var chatID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) (errTx error) {
		chatID, errTx = s.chatRepository.Create(ctx, req.Name)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatUserRepository.Create(ctx, chatID, req.UserNames)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.chatLogRepository.Create(ctx, model.ChatLogInfo{
			ChatID: chatID,
			Log:    createLog,
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatID, nil
}
