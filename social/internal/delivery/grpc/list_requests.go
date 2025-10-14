package social_grpc

import (
	"context"
	"msa_big_tech/social/internal/middleware"
	social "msa_big_tech/social/pkg/proto/v1"
)

// Входящие заявки
func (s *Implementation) ListRequests(ctx context.Context, req *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {

	user_id, err := middleware.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	res, err := s.srv.ListFriendRequests(ctx, user_id)

	if err != nil {
		return nil, err
	}

	a := make([]*social.FriendRequest, len(res))

	for _, val := range res {

		w := &social.FriendRequest{
			RequestId:  val.ID,
			FromUserId: val.FromUserID,
			ToUserId:   val.ToUserID,
			Status:     ConvertStatusToProto(val.Status),
		}

		a = append(a, w)
	}

	return &social.ListRequestsResponse{
		Requests: a,
	}, nil
}
