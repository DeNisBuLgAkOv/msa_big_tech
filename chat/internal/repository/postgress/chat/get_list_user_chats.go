package repo_chat

import (
	"context"
	"msa_big_tech/chat/internal/models"
	"time"
)

func (s *ChatRepository) GetListUserChats(ctx context.Context, userID string) ([]*models.Chat, error) {

	arr := []*models.Chat{
		{
			ID:        "asdasdasdas",
			UserIDs:   []string{"sdasd-asdada-asdasd", "sdasd-asdadasdadwe-asdsadasdsa"},
			Messages:  []models.Message{{ID: 1, Text: "sdada"}},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "asdasdaasdasdasdasdasdsdas",
			UserIDs:   []string{"sdasdasdasd-asdada-asdasd", "sdasd-asdadaasdasdsdadwe-asdsadasdsa"},
			Messages:  []models.Message{{ID: 1, Text: "sdada"}},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return arr, nil

}
