package dto

import "msa_big_tech/chat/internal/models"

type CreateDirectChatDto struct {
	UserID        string
	ParticipantID string
}

type GetChatDto struct {
	UserID string
	ChatID string
}

type ListUserChatsDto struct {
	UserID string
}

type ListChatMembersDto struct {
	UserID string
	ChatID string
}

type SendMessageDto struct {
	UserID string
	ChatID string
	Text   string
}

type ListMessagesDto struct {
	UserID string
	ChatID string
	Limit  int32
	Cursor *string
}

type ListMessagesResponse struct {
	Messages   []*models.Message
	NextCursor *string
}
