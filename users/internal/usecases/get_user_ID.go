package usecases

import (
	"context"
	"fmt"
	"msa_big_tech/users/internal/models"
)

func (s *UseCase) GetProfileByID(ctx context.Context, ID string) (*models.User, error) {

	user, err := s.repo.FetchById(ctx, ID)

	if err != nil {
		return nil, fmt.Errorf("GetProfileByID: FetchById: %w", err)
	}

	if user == nil {
		return nil, models.ErrNotFound
	}

	return user, nil
}
