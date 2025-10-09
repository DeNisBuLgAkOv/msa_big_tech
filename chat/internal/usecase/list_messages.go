package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/chat/internal/usecase/dto"
)

func (s *UseCase) ListMessages(ctx context.Context, req dto.ListMessagesDto) (*dto.ListMessagesResponse, error) {

	messages, nextCursor, err := s.repos_message.ListMessage(ctx, req.ChatID, int32(req.Limit), *req.Cursor)
	if err != nil {
		return nil, fmt.Errorf("ListMessages ListMessages error: %w", err)
	}

	return &dto.ListMessagesResponse{
		Messages:   messages,
		NextCursor: &nextCursor,
	}, nil
}
