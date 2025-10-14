package usecase

import (
	"context"
	"errors"
	"fmt"
	"msa_big_tech/social/internal/models"
	"msa_big_tech/social/internal/usecase/dto"
)

func (s *UseCase) AcceptFriendRequest(ctx context.Context, req dto.ChangeFriendRequestDto) (*models.FriendRequest, error) {

	//Получить заявку из базы
	res, err := s.repo_request.GetFriendRequest(ctx, req.RequestID)
	if err != nil {
		return nil, models.ErrRequestNotFound
	}

	if res.ToUserID != req.UserID {
		return nil, models.ErrPermissionDenied
	}

	//  Проверить что заявка еще PENDING
	if res.Status != models.FriendRequestPending {
		return nil, errors.New("request already processed")
	}

	// 4. Обновить статус заявки на ACCEPTED
	updatedRequest, err := s.repo_request.UpdateFriendRequestStatus(ctx, req.RequestID, models.FriendRequestAccepted)
	if err != nil {
		return nil, err
	}

	// Создаем в таблице друзей пары
	err = s.repo_friends.CreateFriendship(ctx, res.FromUserID, res.ToUserID)

	if err != nil {
		return nil, fmt.Errorf("AcceptFriendRequest CreateFriendship error: %w", err)
	}

	return &models.FriendRequest{
		ID:         updatedRequest.ID,
		FromUserID: updatedRequest.FromUserID,
		ToUserID:   updatedRequest.ToUserID,
		Status:     updatedRequest.Status,
		CreatedAt:  updatedRequest.CreatedAt,
		UpdatedAt:  updatedRequest.UpdatedAt,
	}, nil
}
