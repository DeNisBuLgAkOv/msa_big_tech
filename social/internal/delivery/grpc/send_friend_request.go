package social_grpc

import (
	"context"
	"msa_big_tech/social/internal/middleware"
	"msa_big_tech/social/internal/usecase/dto"
	social "msa_big_tech/social/pkg/proto/v1"
)

// Отправить заявку в друзья
func (s *Implementation) SendFriendRequest(ctx context.Context, req *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	data := &dto.FriendRequestDto{
		ToUserID:   req.ToUserId,
		FromUserID: user_id,
	}

	res, err := s.srv.SendFriendRequest(ctx, *data)
	if err != nil {
		return nil, err
	}

	return &social.SendFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: res.ID,
			Status:    convertStatusToProto(res.Status),
		},
	}, nil
}

func convertStatusToProto(status string) social.Status {
	switch status {
	case "PENDING", "pending":
		return social.Status_STATUS_PENDING
	case "ACCEPTED", "accepted", "APPROVED", "approved":
		return social.Status_STATUS_APPROVED
	case "DECLINED", "declined":
		return social.Status_STATUS_DECLINED
	default:
		return social.Status_STATUS_PENDING
	}
}
