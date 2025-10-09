package token_repo

import (
	"context"
	"time"
)

// CreateToken - аутентификация пользователя и создание токенов
func (r *Repository) CreateToken(ctx context.Context, ID string) (*RefreshToken, error) {

	refresh := "qwertyuiopasdfghjkl"
	expiresAt, _ := time.Parse("2006-01-02 15:04:05", "2025-10-15 14:30:00")

	return &RefreshToken{
		UserID:    "какойто-классный-uuid",
		Token:     refresh,
		ExpiresAt: expiresAt,
	}, nil
}
