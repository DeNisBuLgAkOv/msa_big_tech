package models

import "time"

type Message struct {
	ID        uint64
	Text      string
	ChatID    string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Chat struct {
	ID        string
	UserIDs   []string
	Messages  []Message
	CreatedAt time.Time
	UpdatedAt time.Time
}
