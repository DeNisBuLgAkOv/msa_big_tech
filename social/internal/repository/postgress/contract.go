package repo_postgress

import (
	"context"
	"errors"
	"msa_big_tech/social/internal/models"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrPermissionDenied = errors.New("permission denied")
)

// === Управление заявками в друзья ===
type SocialRequestRepository interface {

	// CreateFriendRequest - создает новую заявку в друзья
	CreateFriendRequest(ctx context.Context, fromUserID, toUserID string) (*models.FriendRequest, error) ////+++++++++

	// GetFriendRequest - получает заявку по ID
	GetFriendRequest(ctx context.Context, requestID string) (*models.FriendRequest, error) ///++++++

	// GetFriendRequestByUsers - получает заявку между двумя пользователями
	// Используется для проверки существующей заявки
	GetFriendRequestByUsers(ctx context.Context, fromUserID, toUserID string) (*models.FriendRequest, error) //++++++

	// ListIncomingFriendRequests - список входящих заявок (где userID - получатель)
	ListIncomingFriendRequests(ctx context.Context, userID string) ([]*models.FriendRequest, error) //+++++++++

	// UpdateFriendRequestStatus - обновляет статус заявки
	UpdateFriendRequestStatus(ctx context.Context, requestID string, status string) (*models.FriendRequest, error) ///+++++

	// DeleteFriendRequest - удаляет заявку (при отмене или очистке)
	DeleteFriendRequest(ctx context.Context, requestID string) error ///++++
}

// === Управление друзьями ===
type SocialFriebdRepository interface {
	// CreateFriendship - создает дружескую связь (после принятия заявки)
	// Создает две записи: (user1, user2) и (user2, user1)
	CreateFriendship(ctx context.Context, userID1, userID2 string) error ////+++++=

	// RemoveFriendship - удаляет дружескую связь
	// Удаляет обе записи: (user1, user2) и (user2, user1)
	RemoveFriendship(ctx context.Context, userID1, userID2 string) error ////+++++++

	// ListFriends - возвращает список друзей пользователя с пагинацией
	ListFriends(ctx context.Context, userID string, limit int, cursor string) ([]string, string, error)
}
