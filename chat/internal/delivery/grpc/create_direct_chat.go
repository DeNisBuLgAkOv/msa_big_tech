package delivery_grpc

import (
	"context"
	"msa_big_tech/chat/internal/middleware"
	"msa_big_tech/chat/internal/usecase/dto"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

// CreateDirectChat создаёт личный чат.
func (s *Implementation) CreateDirectChat(ctx context.Context, req *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	data := dto.CreateDirectChatDto{
		UserID:        user_id,
		ParticipantID: req.ParticipantId,
	}

	chatID, err := s.srv.CreateDirectChat(ctx, data)
	if err != nil {
		return nil, err
	}

	return &chat.CreateDirectChatResponse{
		ChatId: chatID,
	}, nil
}
