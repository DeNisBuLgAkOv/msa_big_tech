package repo_chat

import (
	"context"
	"msa_big_tech/chat/internal/models"
	"time"
)

func (s *ChatRepository) GetChat(ctx context.Context, user1_id, user2_id string) (*models.Chat, error) {

	return &models.Chat{
		ID:        "asdasdasdas",
		UserIDs:   []string{"sdasd-asdada-asdasd", "sdasd-asdadasdadwe-asdsadasdsa"},
		Messages:  []models.Message{{ID: 1, Text: "sdada"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
