package usecases

import (
	"context"
	"fmt"
	"msa_big_tech/users/internal/models"
	"msa_big_tech/users/internal/usecases/dto"
)

func (s *UseCase) CreateUser(ctx context.Context, u *dto.User) (*models.User, error) {

	if user, err := s.repo.FetchByNickname(ctx, u.Nickname); err != nil {
		return nil, fmt.Errorf("CreateUser: FetchByNickname: %w", err)
	} else if user != nil {
		return nil, models.ErrAlreadyExists
	}

	user, err := s.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}
