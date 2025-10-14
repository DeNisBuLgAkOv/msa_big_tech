package social_grpc

import (
	"context"
	"msa_big_tech/social/internal/middleware"
	"msa_big_tech/social/internal/usecase/dto"
	social "msa_big_tech/social/pkg/proto/v1"
)

// Принять заявку
func (s *Implementation) AcceptFriendRequest(ctx context.Context, req *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}
	data := dto.ChangeFriendRequestDto{
		UserID:    user_id,
		RequestID: req.RequestId,
	}

	res, err := s.srv.AcceptFriendRequest(ctx, data)
	if err != nil {
		return nil, err
	}

	return &social.AcceptFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: res.ID,
			Status:    ConvertStatusToProto(res.Status),
		},
	}, nil
}
