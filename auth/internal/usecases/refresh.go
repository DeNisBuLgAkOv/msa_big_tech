package usecases

import (
	"context"
	"errors"
	"msa_big_tech/auth/internal/models"
	token_repo "msa_big_tech/auth/internal/repository/postgress/token"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
	ErrUserNotFound        = errors.New("user not found")
)

func (s *UseCase) Refresh(ctx context.Context, a *usecase_dto.AuthRefresh) (*models.Tokens, error) {

	var newRefreshToken *token_repo.RefreshToken

	// Получаем refresh token из базы
	result, err := s.token.GetByToken(ctx, a.RefreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	// Проверяем не истек ли токен
	if time.Now().After(result.ExpiresAt) {
		// Удаляем просроченный токен
		err = s.token.DeleteByToken(ctx, result.Token)
		if err != nil {
			return nil, err
		}
		newRefreshToken, err = s.token.CreateToken(ctx, a.ID)
		if err != nil {
			return nil, err
		}

		return nil, ErrRefreshTokenExpired
	}

	//генерация accessToken
	accessToken, err := GenerateAccessToken(a.ID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	return &models.Tokens{
		UserID:       a.ID,
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken.Token,
	}, nil
}

// GenerateAccessToken создает JWT access token
func GenerateAccessToken(userID string) (string, error) {
	secret := "моковыйСекрет" // TODO: вынести в дальнейшем
	expiresAt := time.Now().Add(15 * time.Minute)

	claims := jwt.StandardClaims{
		ExpiresAt: expiresAt.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "auth_service",
		Subject:   userID,
		Id:        uuid.New().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// generateRefreshToken создает случайный refresh token
func GenerateRefreshToken() (string, time.Time, error) {
	// Просто UUID
	token := uuid.New().String()
	expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 дней

	return token, expiresAt, nil
}
