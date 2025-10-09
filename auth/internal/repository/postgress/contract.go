package postgress

import (
	"context"
	"msa_big_tech/auth/internal/models"
	token_repo "msa_big_tech/auth/internal/repository/postgress/token"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, req *models.User) (string, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, ID string) (*token_repo.RefreshToken, error)
	GetByToken(ctx context.Context, refreshToken string) (*token_repo.RefreshToken, error)
	DeleteByToken(ctx context.Context, refreshToken string) error
}
