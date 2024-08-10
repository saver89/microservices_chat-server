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

// ChatInfo is the chat info model
type ChatInfo struct {
	Name      string
	UserNames []string
}

type Message struct {
	ID        int64
	Info      MessageInfo
	CreatedAt time.Time
}

type MessageInfo struct {
	FromUser string
	Text     string
	ChatID   int64
	SentAt   time.Time
}
