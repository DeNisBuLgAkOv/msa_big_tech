package user_repo

import (
	"context"
	"msa_big_tech/auth/internal/models"
)

func (s *Repository) Create(ctx context.Context, req *models.User) (string, error) {

	return "какойто-классный-uuid", nil
}
