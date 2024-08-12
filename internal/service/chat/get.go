package chat

import (
	"context"

	"github.com/saver89/microservices_chat-server/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Chat, error) {
	var chat *model.Chat

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) (errTx error) {
		chat, errTx = s.chatRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		userNames, errTx := s.chatUserRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}
		chat.Info.UserNames = userNames

		_, errTx = s.chatLogRepository.Create(ctx, model.ChatLogInfo{
			ChatID: id,
			Log:    getLog,
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return chat, nil
}
