package repo_request

import (
	"context"
	"msa_big_tech/social/internal/models"
	"time"
)

func (s *RequestsRepository) ListIncomingFriendRequests(ctx context.Context, userID string) ([]*models.FriendRequest, error) {

	return []*models.FriendRequest{
		{
			ID:         "req_1",
			FromUserID: "user_1",
			ToUserID:   userID,
			Status:     "PENDING",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         "req_2",
			FromUserID: "user_2",
			ToUserID:   userID,
			Status:     "PENDING",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}, nil
}
