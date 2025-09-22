package transport

import (
	"context"
	"log"
	auth "msa_big_tech/auth/pkg/proto/v1"
)

type Implementation struct {
	auth.UnimplementedAuthServiceServer
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

func (s *Implementation) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Println(req.Email)
	return &auth.RegisterResponse{
		UserId: "sadasdadsdsad",
	}, nil
}

func (s *Implementation) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{}, nil
}
func (s *Implementation) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return &auth.RefreshResponse{}, nil
}
