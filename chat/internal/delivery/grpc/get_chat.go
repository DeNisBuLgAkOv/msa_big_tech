package delivery_grpc

import (
	"context"
	"msa_big_tech/chat/internal/middleware"
	"msa_big_tech/chat/internal/usecase/dto"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

// GetChat возвращает информацию о чате.
func (s *Implementation) GetChat(ctx context.Context, req *chat.GetChatRequest) (*chat.GetChatResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	data := dto.GetChatDto{
		UserID: user_id,
		ChatID: req.ChatId,
	}

	res, err := s.srv.GetChat(ctx, data)
	if err != nil {
		return nil, err
	}
	message := make([]*chat.Message, len(res.Messages))

	for _, val := range res.Messages {
		data := chat.Message{
			ChatId:    val.ChatID,
			MessageId: val.ID,
			Text:      val.Text,
			UserId:    val.UserID,
		}
		message = append(message, &data)
	}

	conver := &chat.Chat{
		ChatId:   res.ID,
		UserIds:  res.UserIDs,
		Messages: message,
	}

	return &chat.GetChatResponse{
		Chat: conver,
	}, nil
}
