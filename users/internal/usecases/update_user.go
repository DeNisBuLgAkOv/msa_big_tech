package usecases

import (
	"context"
	"fmt"
	"msa_big_tech/users/internal/models"
	"msa_big_tech/users/internal/usecases/dto"
)

func (s *UseCase) UpdateUser(ctx context.Context, u *dto.User) (*models.User, error) {

	user, err := s.repo.FetchById(ctx, u.ID)

	if err != nil {
		return nil, fmt.Errorf("UpdateUser: FetchById: %w", err)
	}
	if user == nil {
		return nil, models.ErrNotFound
	}

	user, err = s.repo.Update(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser: Update: %w", err)
	}

	return user, nil
}
