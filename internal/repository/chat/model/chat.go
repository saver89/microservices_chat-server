package model

import (
	"time"
)

// Chat is the chat model
type Chat struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
