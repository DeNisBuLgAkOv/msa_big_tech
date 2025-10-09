package user_repo

import (
	"context"
	"msa_big_tech/auth/internal/models"
)

func (s *Repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {

	return &models.User{
		ID:       "какойто-классный-uuid",
		Email:    "test@mail.ru",
		Password: "$2a$10$uYR2vJHpRQUEDqJ7C6S7JOqY.1o/jeqhtpH3WBT1trQm39gD3rQtq", //хэш

	}, nil
	// return nil, nil
}
