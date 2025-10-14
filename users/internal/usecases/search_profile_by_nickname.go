package usecases

import (
	"context"
	"fmt"
	"msa_big_tech/users/internal/models"
	"msa_big_tech/users/internal/usecases/dto"
)

func (s *UseCase) SearchByNickname(ctx context.Context, req dto.SearchByNicknameRequest) ([]*models.User, error) {
	users, err := s.repo.SearchByNickname(ctx, req.Query, req.Limit)
	if err != nil {
		return nil, fmt.Errorf("SearchByNickname: SearchByNickname error: %w", err)
	}

	return users, nil
}
