package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/chat/internal/models"
	"msa_big_tech/chat/internal/usecase/dto"
)

func (s *UseCase) CreateDirectChat(ctx context.Context, req dto.CreateDirectChatDto) (string, error) {

	// Проверяем нет ли уже чата между этими пользователями
	existingChat, err := s.repo_chat.GetChat(ctx, req.UserID, req.ParticipantID)
	if err == nil && existingChat != nil {
		return "", models.ErrChatAlreadyExists
	}

	// Создаем новый чат
	chat, err := s.repo_chat.Create(ctx, req.UserID, req.ParticipantID)
	if err != nil {
		return "", fmt.Errorf("CreateDirectChat Create error: %w", err)
	}

	return chat, nil

}
