package postgress

import (
	"context"
	"msa_big_tech/chat/internal/models"
	"msa_big_tech/chat/internal/usecase/dto"
)

type ChatRepository interface {
	Create(ctx context.Context, user1_id, user2_id string) (string, error)
	GetChat(ctx context.Context, user1_id, user2_id string) (*models.Chat, error)
	GetListUserChats(ctx context.Context, userID string) ([]*models.Chat, error)
	GetUsersByChatID(ctx context.Context, chatID string) ([]string, error)
}

type MessageRepository interface {
	SendMessage(ctx context.Context, req *dto.SendMessageDto) (*models.Message, error)
	ListMessage(ctx context.Context, ChatID, limit, cursor string) ([]*models.Message, error)
}
