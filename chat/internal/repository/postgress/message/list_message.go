package repo_message

import (
	"context"
	"msa_big_tech/chat/internal/models"
)

func (s *MessageRepository) ListMessage(ctx context.Context, ChatID string, limit int32, cursor string) ([]*models.Message, string, error) {
	return []*models.Message{
		&models.Message{ID: 2, Text: "sdsdadas", ChatID: "asdasdassad"},
	}, "adasdas", nil
}
