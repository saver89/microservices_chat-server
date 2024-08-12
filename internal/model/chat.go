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

// Message is the message model
type Message struct {
	ID        int64
	Info      MessageInfo
	CreatedAt time.Time
}

// MessageInfo is the message info model
type MessageInfo struct {
	FromUser string
	Text     string
	ChatID   int64
	SentAt   time.Time
}
