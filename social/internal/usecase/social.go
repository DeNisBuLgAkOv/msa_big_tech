package usecase

import (
	repos_friends "msa_big_tech/social/internal/repository/postgress/friends"
	repo_request "msa_big_tech/social/internal/repository/postgress/request"
)

type UseCase struct {
	repo_friends repos_friends.FriendsRepository
	repo_request repo_request.RequestsRepository
}

func NewUsecase(repo_friends repos_friends.FriendsRepository,
	repo_request repo_request.RequestsRepository) *UseCase {
	return &UseCase{repo_friends: repo_friends, repo_request: repo_request}

}
