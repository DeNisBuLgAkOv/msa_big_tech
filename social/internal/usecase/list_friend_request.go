package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/social/internal/models"
)

// ListFriendRequests получение списка заявок на друзья
func (s *UseCase) ListFriendRequests(ctx context.Context, toUserID string) ([]*models.FriendRequest, error) {

	res, err := s.repo_request.ListIncomingFriendRequests(ctx, toUserID)
	if err != nil {
		return nil, fmt.Errorf("ListFriendRequests ListIncomingFriendRequests error: %w", err)
	}

	return res, nil

}
