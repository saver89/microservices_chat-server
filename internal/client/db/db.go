package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler is the handler function
type Handler func(ctx context.Context) error

// Client is the client interface
type Client interface {
	DB() DB
	Close() error
}

// TxManager is the transaction manager interface
type TxManager interface {
	ReadCommitted(ctx context.Context, handler Handler) error
}

// Query is the query struct
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor is the interface for transactions
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer is the interface for executing SQL queries
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer is the interface for executing named queries
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer is the interface for executing queries
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger is the interface for pinging the database
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB is the database interface
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
