package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/saver89/microservices_chat-server/internal/client/db"
	"github.com/saver89/microservices_chat-server/internal/repository"
)

const (
	tableName = "chat_users"

	idColumn       = "id"
	chatIDColumn   = "chat_id"
	userNameColumn = "user_name"
)

type repo struct {
	db db.Client
}

// NewChatRepository creates a new chat repository
func NewChatUserRepository(db db.Client) repository.ChatUserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, chatID int64, userNames []string) (err error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userNameColumn)

	for _, userName := range userNames {
		builder = builder.Values(chatID, userName)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_user_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(ctx context.Context, chatID int64) ([]string, error) {
	builder := sq.Select(userNameColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{chatIDColumn: chatID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_user_repository.Get",
		QueryRaw: query,
	}

	var userNames []string
	err = r.db.DB().ScanAllContext(ctx, &userNames, q, args...)
	if err != nil {
		return nil, err
	}

	return userNames, nil
}

func (r *repo) Delete(ctx context.Context, chatID int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: chatID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
