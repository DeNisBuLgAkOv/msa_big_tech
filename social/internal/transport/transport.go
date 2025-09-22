package transport

import (
	"context"
	social "msa_big_tech/social/pkg/proto/v1"
)

type Implementation struct {
	social.UnimplementedSocialServiceServer
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

func (s *Implementation) SendFriendRequest(ctx context.Context, req *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {
	return &social.SendFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: "request_123",
			Status:    social.Status_STATUS_PENDING,
		},
	}, nil
}

func (s *Implementation) ListRequests(ctx context.Context, req *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {
	return &social.ListRequestsResponse{
		Requests: []*social.FriendRequest{
			{
				RequestId: "request_123",
				Status:    social.Status_STATUS_PENDING,
			},
			{
				RequestId: "request_124",
				Status:    social.Status_STATUS_PENDING,
			},
		},
	}, nil
}

func (s *Implementation) AcceptFriendRequest(ctx context.Context, req *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	return &social.AcceptFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: req.RequestId,
			Status:    social.Status_STATUS_APPROVED,
		},
	}, nil
}

func (s *Implementation) DeclineFriendRequest(ctx context.Context, req *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {
	return &social.DeclineFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: "asdasds",
			Status:    social.Status_STATUS_DECLINED,
		},
	}, nil
}

func (s *Implementation) RemoveFriend(ctx context.Context, req *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {

	return &social.RemoveFriendResponse{}, nil
}
func (s *Implementation) ListFriends(ctx context.Context, req *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	friendIds := []string{"friend_456", "friend_789"}
	nextCursor := ""
	if len(friendIds) == int(req.Limit) {
		nextCursor = friendIds[len(friendIds)-1]
	}

	return &social.ListFriendsResponse{
		FriendUserIds: friendIds,
		NextCursor:    nextCursor,
	}, nil
}
