package usecases

import (
	user_repo "msa_big_tech/auth/internal/repository/postgress/auth"
	token_repo "msa_big_tech/auth/internal/repository/postgress/token"
)

type UseCase struct {
	userRepo *user_repo.Repository
	token    *token_repo.Repository
}

func NewUseCase(userRepo *user_repo.Repository,
	token *token_repo.Repository) *UseCase {
	return &UseCase{userRepo: userRepo, token: token}
}
