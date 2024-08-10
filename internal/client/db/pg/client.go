package pg

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/saver89/microservices_chat-server/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

// New creates a new pg client
func New(ctx context.Context, dsn string, log *slog.Logger) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to pg: %v", err)
	}

	return &pgClient{
		masterDBC: NewDB(dbc, log),
	}, nil
}

// DB returns the database
func (p pgClient) DB() db.DB {
	return p.masterDBC
}

// Close closes the client
func (p pgClient) Close() error {
	if p.masterDBC == nil {
		p.masterDBC.Close()
	}

	return nil
}
