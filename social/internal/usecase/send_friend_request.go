package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/social/internal/models"
	"msa_big_tech/social/internal/usecase/dto"
)

func (s *UseCase) SendFriendRequest(ctx context.Context, req dto.FriendRequestDto) (*models.FriendRequest, error) {

	res, err := s.repo_request.GetFriendRequestByUsers(ctx, req.FromUserID, req.ToUserID)
	if err != nil {
		return nil, fmt.Errorf("SendFriendRequest GetFriendRequestByUsers error: %w", err)
	} else if res == nil {
		return nil, models.ErrRequestExists
	}

	res, err = s.repo_request.CreateFriendRequest(ctx, req.FromUserID, req.ToUserID)
	if err != nil {
		return nil, fmt.Errorf("SendFriendRequest CreateFriendRequest error: %w", err)
	}

	return &models.FriendRequest{
		ID:         res.ID,
		FromUserID: res.FromUserID,
		ToUserID:   res.ToUserID,
		Status:     res.Status,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}, nil
}
