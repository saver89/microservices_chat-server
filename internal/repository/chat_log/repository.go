package chat_log

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/saver89/microservices_chat-server/internal/client/db"
	"github.com/saver89/microservices_chat-server/internal/model"
	"github.com/saver89/microservices_chat-server/internal/repository"
)

const (
	tableName = "chat_logs"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	logColumn       = "log"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewChatLogRepository creates a new chat log repository
func NewChatLogRepository(db db.Client) repository.ChatLogRepository {
	return &repo{
		db: db,
	}
}

// Create creates a new chat log
func (r *repo) Create(ctx context.Context, log model.ChatLogInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, logColumn).
		Values(log.ChatID, log.Log).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_log_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
