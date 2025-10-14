package usecase

import (
	"context"
	"msa_big_tech/social/internal/models"
	"msa_big_tech/social/internal/usecase/dto"
)

type SocialUsecases interface {
	// SendFriendRequest отправка заявки на друзья
	SendFriendRequest(ctx context.Context, req dto.FriendRequestDto) (*models.FriendRequest, error) /////+
	// ListFriendRequests получение списка заявок на друзья
	ListFriendRequests(ctx context.Context, toUserID string) ([]*models.FriendRequest, error) ///+++
	// AcceptFriendRequest принятие заявки на друзья
	AcceptFriendRequest(ctx context.Context, req dto.ChangeFriendRequestDto) (*models.FriendRequest, error) ////++++
	// DeclineFriendRequest отказ от заявки на друзья
	DeclineFriendRequest(ctx context.Context, req dto.ChangeFriendRequestDto) (*models.FriendRequest, error) ///++++++
	// RemoveFriend удаление друга
	RemoveFriend(ctx context.Context, req dto.FriendRequestDto) error ///++++
	// ListFriends получение списка друзей
	ListFriends(ctx context.Context, req dto.ListFriendsDto) (*dto.ListFriendsResponse, error)
}
