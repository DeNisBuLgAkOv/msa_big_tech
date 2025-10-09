package delivery_grpc

import (
	"context"
	"msa_big_tech/chat/internal/middleware"
	"msa_big_tech/chat/internal/usecase/dto"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

// SendMessage отправляет сообщение в чат.
func (s *Implementation) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	data := dto.SendMessageDto{
		UserID: user_id,
		ChatID: req.ChatId,
		Text:   req.Text,
	}

	res, err := s.srv.SendMessage(ctx, data)
	if err != nil {
		return nil, err
	}

	return &chat.SendMessageResponse{
		Message: &chat.Message{
			MessageId: res.ID,
			Text:      res.Text,
			ChatId:    res.ChatID,
			UserId:    user_id,
		},
	}, nil
}
