package users_repo

import (
	"context"
	"msa_big_tech/users/internal/models"
	"msa_big_tech/users/internal/usecases/dto"
)

type UsersRepository interface {
	Create(ctx context.Context, user *dto.User) (*models.User, error)
	Update(ctx context.Context, user *dto.User) (*models.User, error)
	FetchById(ctx context.Context, id string) (*models.User, error)
	FetchByNickname(ctx context.Context, nickname string) (*models.User, error)
	SearchByNickname(ctx context.Context, query string, limit int64) ([]*models.User, error)
}

// Create создает пользователя
func (r *Repository) Create(ctx context.Context, user *dto.User) (*models.User, error) {
	return &models.User{
		ID:       "user-123",
		Nickname: "new_user",
		Bio:      "Новый пользователь",
		Avatar:   "https://example.com/avatar1.jpg",
	}, nil
}

// Обновить пользователя
func (r *Repository) Update(ctx context.Context, user *dto.User) (*models.User, error) {
	return &models.User{
		ID:       "user-456",
		Nickname: "updated_user",
		Bio:      "Обновленная информация",
		Avatar:   "https://example.com/avatar2.jpg",
	}, nil
}

// Найти пользователя по ID
func (r *Repository) FetchById(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
	// return &models.User{
	// 	ID:       "user-789",
	// 	Nickname: "john_doe",
	// 	Bio:      "Разработчик из Москвы",
	// 	Avatar:   "https://example.com/avatar3.jpg",
	// }, nil
}

// Найти пользователя по нику
func (r *Repository) FetchByNickname(ctx context.Context, nickname string) (*models.User, error) {
	return nil, nil
	// return &models.User{
	// 	ID:       "user-999",
	// 	Nickname: "alice_smith",
	// 	Bio:      "Дизайнер интерфейсов",
	// 	Avatar:   "https://example.com/avatar4.jpg",
	// }, nil
}

func (r *Repository) SearchByNickname(ctx context.Context, query string, limit int64) ([]*models.User, error) {
	return []*models.User{
		{
			ID:       "user-111",
			Nickname: "john_doe",
			Bio:      "Разработчик из Москвы",
			Avatar:   "https://example.com/avatar5.jpg",
		},
		{
			ID:       "user-222",
			Nickname: "john_smith",
			Bio:      "Бэкенд разработчик",
			Avatar:   "https://example.com/avatar6.jpg",
		},
		{
			ID:       "user-333",
			Nickname: "johnny_depp",
			Bio:      "Актер",
			Avatar:   "https://example.com/avatar7.jpg",
		},
	}, nil
}
