package users_grpc

import (
	"msa_big_tech/users/internal/usecases"
	users "msa_big_tech/users/pkg/proto/v1"
)

type Implementation struct {
	users.UnimplementedUsersServiceServer
	usecases usecases.UsersUsecases
}

func NewImplementation(usecases usecases.UsersUsecases) *Implementation {

	return &Implementation{
		usecases: usecases,
	}
}
