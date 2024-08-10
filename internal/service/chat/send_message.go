package chat

import (
	"context"

	"github.com/saver89/microservices_chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, req *model.MessageInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) (errTx error) {
		errTx = s.messageRepository.SendMessage(ctx, req)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.chatLogRepository.Create(ctx, model.ChatLogInfo{
			ChatID: req.ChatID,
			Log:    sendMessageLog,
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
