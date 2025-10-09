package usecase

import (
	"context"
	"msa_big_tech/chat/internal/models"
	"msa_big_tech/chat/internal/usecase/dto"
)

type ChatUsecase interface {
	// CreateDirectChat создание личного чата
	CreateDirectChat(ctx context.Context, req dto.CreateDirectChatDto) (string, error) //+++
	// GetChat получение информации о чате
	GetChat(ctx context.Context, req dto.GetChatDto) (*models.Chat, error)
	// GetListUserChats получение списка чатов пользователя
	GetListUserChats(ctx context.Context, userID string) ([]*models.Chat, error)
	// ListChatMembers получение участников чата
	ListChatMembers(ctx context.Context, chatID string) ([]string, error)
	// SendMessage отправка сообщения
	SendMessage(ctx context.Context, req dto.SendMessageDto) (*models.Message, error)
	// ListMessages получение истории сообщений
	ListMessages(ctx context.Context, req dto.ListMessagesDto) (*dto.ListMessagesResponse, error)
	// StreamMessages серверный стрим новых сообщений (пока непонятно как, позже реализую)
	StreamMessages(ctx context.Context)
}
