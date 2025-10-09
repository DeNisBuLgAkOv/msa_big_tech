package social_grpc

import (
	"context"
	"msa_big_tech/social/internal/middleware"
	"msa_big_tech/social/internal/usecase/dto"
	social "msa_big_tech/social/pkg/proto/v1"
)

// Список друзей
func (s *Implementation) ListFriends(ctx context.Context, req *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	data := dto.ListFriendsDto{
		UsersID: user_id,
		Limit:   int64(req.Limit),
		Cursor:  &req.Cursor,
	}

	res, err := s.srv.ListFriends(ctx, data)

	if err != nil {
		return nil, err
	}

	return &social.ListFriendsResponse{
		FriendUserIds: res.Friends,
		NextCursor:    *res.NextCursor,
	}, nil
}
