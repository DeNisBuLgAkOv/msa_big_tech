package delivery_grpc

import (
	"context"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

// ListChatMembers возвращает список участников чата.
func (s *Implementation) ListChatMembers(ctx context.Context, req *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {

	res, err := s.srv.ListChatMembers(ctx, req.ChatId)
	if err != nil {
		return nil, err
	}

	return &chat.ListChatMembersResponse{
		UserIds: res,
	}, nil
}
