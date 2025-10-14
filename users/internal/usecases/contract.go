package usecases

import (
	"context"
	"msa_big_tech/users/internal/models"
	"msa_big_tech/users/internal/usecases/dto"
)

type UsersUsecases interface {
	// Создание профиля пользователя
	CreateUser(ctx context.Context, u *dto.User) (*models.User, error)
	// Обновление профиля пользователя
	UpdateUser(ctx context.Context, u *dto.User) (*models.User, error)
	// Получение профиля пользователя по id
	GetProfileByID(ctx context.Context, id string) (*models.User, error)
	// Получение профиля пользователя по nickname
	GetProfileByNickname(ctx context.Context, nickname string) (*models.User, error)
}
