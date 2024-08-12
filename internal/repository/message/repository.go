package message

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/saver89/microservices_chat-server/internal/client/db"
	"github.com/saver89/microservices_chat-server/internal/model"
	"github.com/saver89/microservices_chat-server/internal/repository"
)

const (
	tableName = "messages"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	fromUserColumn  = "from_user"
	textColumn      = "text"
	sentAtColumn    = "sent_at"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewMessageRepository creates a new message repository
func NewMessageRepository(db db.Client) repository.MessageRepository {
	return &repo{
		db: db,
	}
}

// Create creates a new chat log
func (r *repo) SendMessage(ctx context.Context, req *model.MessageInfo) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, fromUserColumn, textColumn, sentAtColumn).
		Values(req.ChatID, req.FromUser, req.Text, req.SentAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_log_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
