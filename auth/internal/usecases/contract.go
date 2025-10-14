package usecases

import (
	"context"
	"msa_big_tech/auth/internal/models"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
)

type AuthUseCases interface {
	// Регистрация
	Register(ctx context.Context, a *usecase_dto.Register) (string, error)
	// Логин
	Login(ctx context.Context, a *usecase_dto.Login) (*models.Tokens, error)
	// Обновление токенов
	Refresh(ctx context.Context, a *usecase_dto.AuthRefresh) (*models.Tokens, error)
}
