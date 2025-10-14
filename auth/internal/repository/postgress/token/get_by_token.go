package token_repo

import (
	"context"
	"time"
)

type RefreshToken struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

func (s *Repository) GetByToken(ctx context.Context, refreshToken string) (*RefreshToken, error) {
	expiresAt, _ := time.Parse("2006-01-02 15:04:05", "2025-10-15 14:30:00")
	return &RefreshToken{
		UserID:    "какойто-классный-uuid",
		Token:     "qwertyuiopasdfghjkl",
		ExpiresAt: expiresAt,
	}, nil
}
