package delivery_grpc

import (
	"context"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

func (s *Implementation) ListUserChats(ctx context.Context, req *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {

	res, err := s.srv.GetListUserChats(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	coverArr := make([]*chat.Chat, len(res))

	for _, val := range res {

		arrMessage := make([]*chat.Message, len(val.Messages))

		for _, m := range val.Messages {
			resMes := chat.Message{
				MessageId: m.ID,
				Text:      m.Text,
				ChatId:    m.ChatID,
				UserId:    m.UserID,
			}

			arrMessage = append(arrMessage, &resMes)
		}

		data := &chat.Chat{
			ChatId:   val.ID,
			UserIds:  val.UserIDs,
			Messages: arrMessage,
		}

		coverArr = append(coverArr, data)

	}

	return &chat.ListUserChatsResponse{
		Chats: coverArr,
	}, nil
}
