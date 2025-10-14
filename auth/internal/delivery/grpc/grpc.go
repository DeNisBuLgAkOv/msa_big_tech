package auth_grpc

import (
	"msa_big_tech/auth/internal/usecases"
	auth "msa_big_tech/auth/pkg/proto/v1"
)

type Implementation struct {
	auth.UnimplementedAuthServiceServer
	usecases *usecases.UseCase
}

func NewImplementation(usecases *usecases.UseCase) *Implementation {

	return &Implementation{
		usecases: usecases,
	}
}
