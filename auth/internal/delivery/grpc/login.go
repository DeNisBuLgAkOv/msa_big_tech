package auth_grpc

import (
	"context"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
	auth "msa_big_tech/auth/pkg/proto/v1"
)

func (s *Implementation) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {

	data := &usecase_dto.Login{
		Email:    req.GetEmail(),
		Password: req.Password,
	}

	res, err := s.usecases.Login(ctx, data)

	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		UserId:       res.UserID,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
