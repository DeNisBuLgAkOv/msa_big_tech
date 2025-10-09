package social_grpc

import (
	"context"
	"msa_big_tech/social/internal/middleware"
	"msa_big_tech/social/internal/usecase/dto"
	social "msa_big_tech/social/pkg/proto/v1"
)

// Удалить пользователя из друзей
func (s *Implementation) RemoveFriend(ctx context.Context, req *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	data := &dto.FriendRequestDto{
		ToUserID:   req.FriendUserId,
		FromUserID: user_id,
	}

	err = s.srv.RemoveFriend(ctx, *data)

	if err != nil {
		return nil, err
	}

	return &social.RemoveFriendResponse{}, nil
}
