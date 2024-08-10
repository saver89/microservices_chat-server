package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/saver89/microservices_chat-server/internal/client/db"
	"github.com/saver89/microservices_chat-server/internal/model"
	"github.com/saver89/microservices_chat-server/internal/repository"
	"github.com/saver89/microservices_chat-server/internal/repository/chat/converter"
	repoChatModel "github.com/saver89/microservices_chat-server/internal/repository/chat/model"
)

const (
	tableName = "chats"

	idColumn        = "id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewChatRepository creates a new chat repository
func NewChatRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, name string) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(nameColumn).
		Values(name).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	builder := sq.Select(idColumn, createdAtColumn, nameColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	chat := repoChatModel.Chat{}
	err = r.db.DB().ScanOneContext(ctx, &chat, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(&chat), nil
}
