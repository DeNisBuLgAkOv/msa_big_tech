package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/social/internal/usecase/dto"
)

func (s *UseCase) RemoveFriend(ctx context.Context, req dto.FriendRequestDto) error {

	//Удаляем из таблицы друзей пары
	err := s.repo_friends.RemoveFriendship(ctx, req.FromUserID, req.ToUserID)
	if err != nil {
		return fmt.Errorf("RemoveFriend RemoveFriendship error: %w", err)
	}

	// Получаем заявку из бд
	res, err := s.repo_request.GetFriendRequestByUsers(ctx, req.FromUserID, req.ToUserID)
	if err != nil {
		return fmt.Errorf("RemoveFriend GetFriendRequestByUsers error: %w", err)
	}

	// Удаляем заявку
	err = s.repo_request.DeleteFriendRequest(ctx, res.ID)
	if err != nil {
		return fmt.Errorf("RemoveFriend DeleteFriendRequest error: %w", err)
	}

	return nil
}
