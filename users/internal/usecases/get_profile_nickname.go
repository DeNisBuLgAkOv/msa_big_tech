package usecases

import (
	"context"
	"fmt"
	"msa_big_tech/users/internal/models"
)

func (s *UseCase) GetProfileByNickname(ctx context.Context, nickname string) (*models.User, error) {
	user, err := s.repo.FetchByNickname(ctx, nickname)
	if err != nil {
		return nil, fmt.Errorf("GetProfileByNickname: FetchByNickname error: %w", err)
	}
	if user == nil {
		return nil, models.ErrNotFound
	}

	return user, nil
}
