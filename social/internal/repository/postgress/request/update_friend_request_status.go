package repo_request

import (
	"context"
	"msa_big_tech/social/internal/models"
	"time"
)

func (s *RequestsRepository) UpdateFriendRequestStatus(ctx context.Context, requestID string, status string) (*models.FriendRequest, error) {

	return &models.FriendRequest{
		ID:         requestID,
		FromUserID: "user_from",
		ToUserID:   "user_to",
		Status:     status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil

}
