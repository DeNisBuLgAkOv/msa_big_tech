package usecase

import (
	"context"
	"fmt"
)

func (s *UseCase) ListChatMembers(ctx context.Context, chatID string) ([]string, error) {

	res, err := s.repo_chat.GetUsersByChatID(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("ListChatMembers GetUsersByChatID error: %w", err)
	}

	return res, nil
}
