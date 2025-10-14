package usecases

import (
	"context"
	"msa_big_tech/auth/internal/models"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
)

func (s *UseCase) Login(ctx context.Context, a *usecase_dto.Login) (*models.Tokens, error) {

	// Получение пользователя
	res, err := s.userRepo.GetByEmail(ctx, a.Email)
	// Пользователь не найден
	if res == nil {
		return nil, models.ErrNotExistedUser
	}
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    res.Email,
		Password: res.Password,
	}

	// Проверка на совподение пароля
	checkPassword := user.CheckPassword(a.Password)

	if checkPassword == false {
		return nil, models.ErrInvalidPassword
	}

	// создание refreshToken
	refreshToken, err := s.token.CreateToken(ctx, res.ID)
	if err != nil {
		return nil, err
	}

	// создание accessToken
	accessToken, err := GenerateAccessToken(res.ID)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{
		UserID:       res.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Token,
	}, nil

}
