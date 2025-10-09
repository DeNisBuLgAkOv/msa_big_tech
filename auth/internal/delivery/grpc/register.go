package auth_grpc

import (
	"context"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
	auth "msa_big_tech/auth/pkg/proto/v1"
)

func (s *Implementation) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	data := &usecase_dto.Register{
		Email:    req.Email,
		Password: req.Password,
	}

	ID, err := s.usecases.Register(ctx, data)

	if err != nil {
		return nil, err
	}
	return &auth.RegisterResponse{
		UserId: ID,
	}, nil
}
