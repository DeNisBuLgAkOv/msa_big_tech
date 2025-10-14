package repo_request

import (
	"context"
	"msa_big_tech/social/internal/models"
	"time"
)

func (s *RequestsRepository) GetFriendRequestByUsers(ctx context.Context, fromUserID, toUserID string) (*models.FriendRequest, error) {
	return &models.FriendRequest{
		ID:         "req_" + fromUserID + "_" + toUserID,
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Status:     "PENDING",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
