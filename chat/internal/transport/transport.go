package transport

import (
	"context"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

type Implementation struct {
	chat.UnimplementedChatServiceServer
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

// CreateDirectChat создаёт личный чат.
func (s *Implementation) CreateDirectChat(ctx context.Context, req *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {

	return &chat.CreateDirectChatResponse{
		ChatId: "chat_123",
	}, nil
}

// GetChat возвращает информацию о чате.
func (s *Implementation) GetChat(ctx context.Context, req *chat.GetChatRequest) (*chat.GetChatResponse, error) {

	return &chat.GetChatResponse{
		Chat: &chat.Chat{
			ChatId: req.ChatId,
			Name:   "Direct Chat",
		},
	}, nil
}

// ListUserChats возвращает список чатов пользователя.
func (s *Implementation) ListUserChats(ctx context.Context, req *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {

	return &chat.ListUserChatsResponse{
		Chats: []*chat.Chat{
			{ChatId: "chat_123", Name: "Direct Chat with Friend1"},
			{ChatId: "chat_124", Name: "Direct Chat with Friend2"},
		},
	}, nil
}

// ListChatMembers возвращает список участников чата.
func (s *Implementation) ListChatMembers(ctx context.Context, req *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {

	return &chat.ListChatMembersResponse{
		UserIds: []string{"user_456", "user_789"},
	}, nil
}

// SendMessage отправляет сообщение в чат.
func (s *Implementation) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

	return &chat.SendMessageResponse{
		Message: &chat.Message{
			MessageId: 1,
			Text:      req.Text,
		},
	}, nil
}

// ListMessages возвращает историю сообщений чата.
func (s *Implementation) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {

	return &chat.ListMessagesResponse{
		Messages: []*chat.Message{
			{MessageId: 1, Text: "Hello, how are you?"},
			{MessageId: 2, Text: "I'm good, thanks!"},
		},
	}, nil
}

// StreamMessages реализует серверный стрим новых сообщений.
func (s *Implementation) StreamMessages(ctx context.Context, req *chat.StreamMessagesRequest) (*chat.StreamMessagesResponse, error) {

	return &chat.StreamMessagesResponse{
		Stream: &chat.Message{
			MessageId: 3,
			Text:      "New message in stream",
		},
	}, nil
}
