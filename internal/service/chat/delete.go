package chat

import (
	"context"

	"github.com/saver89/microservices_chat-server/internal/model"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) (errTx error) {
		errTx = s.chatUserRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.chatLogRepository.Create(ctx, model.ChatLogInfo{
			ChatID: id,
			Log:    deleteLog,
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
