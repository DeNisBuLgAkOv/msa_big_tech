package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/chat/internal/models"
)

func (s *UseCase) GetListUserChats(ctx context.Context, userID string) ([]*models.Chat, error) {

	res, err := s.repo_chat.GetListUserChats(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("GetListUserChats GetListUserChats error: %w", err)
	}

	return res, nil

}
