package usecases

import users_repo "msa_big_tech/users/internal/repository/postgress"

type UseCase struct {
	repo users_repo.UsersRepository
}

func NewUseCase(repo users_repo.UsersRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
