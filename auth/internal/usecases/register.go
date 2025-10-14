package usecases

import (
	"context"
	"msa_big_tech/auth/internal/models"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
)

func (s *UseCase) Register(ctx context.Context, a *usecase_dto.Register) (string, error) {

	creds := &models.User{
		Email:    a.Email,
		Password: a.Password,
	}

	// Проверка на существование пользователя
	user, err := s.userRepo.GetByEmail(ctx, a.Email)
	if err != nil {
		return "", err
	} else if user != nil {
		return "", models.ErrUserAlreadyExists
	}

	// хэширование пароля
	err = creds.HashPassword()
	if err != nil {
		return "", err
	}

	// создание пользователя
	ID, err := s.userRepo.Create(ctx, creds)

	return ID, nil

}
