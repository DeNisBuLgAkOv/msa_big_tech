package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/chat/internal/models"
	"msa_big_tech/chat/internal/usecase/dto"
)

func (s *UseCase) SendMessage(ctx context.Context, req dto.SendMessageDto) (*models.Message, error) {

	res, err := s.repos_message.SendMessage(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("SendMessage SendMessage error: %w", err)
	}

	return res, nil
}
