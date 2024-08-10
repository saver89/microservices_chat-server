package model

import (
	"time"
)

// Chat is the chat model
type Chat struct {
	ID        int64
	Info      ChatInfo
	CreatedAt time.Time
}

type ChatInfo struct {
	Name      string
	UserNames []string
}
