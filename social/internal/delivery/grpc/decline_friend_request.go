package social_grpc

import (
	"context"
	"msa_big_tech/social/internal/middleware"
	"msa_big_tech/social/internal/usecase/dto"
	social "msa_big_tech/social/pkg/proto/v1"
)

// Отклонить заявку
func (s *Implementation) DeclineFriendRequest(ctx context.Context, req *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}
	data := dto.ChangeFriendRequestDto{
		UserID:    user_id,
		RequestID: req.RequestId,
	}

	res, err := s.srv.DeclineFriendRequest(ctx, data)

	return &social.DeclineFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: res.ID,
			Status:    ConvertStatusToProto(res.Status),
		},
	}, nil
}
