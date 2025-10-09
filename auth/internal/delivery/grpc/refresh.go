package auth_grpc

import (
	"context"
	usecase_dto "msa_big_tech/auth/internal/usecases/dto"
	auth "msa_big_tech/auth/pkg/proto/v1"
)

func (s *Implementation) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {

	data := &usecase_dto.AuthRefresh{
		RefreshToken: req.RefreshToken,
	}

	tokens, err := s.usecases.Refresh(ctx, data)

	if err != nil {
		return nil, err
	}

	return &auth.RefreshResponse{
		UserId:       tokens.UserID,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
