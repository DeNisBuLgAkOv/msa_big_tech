package social_grpc

import (
	"msa_big_tech/social/internal/usecase"
	social "msa_big_tech/social/pkg/proto/v1"
)

type Implementation struct {
	social.UnimplementedSocialServiceServer
	srv usecase.SocialUsecases
}

func NewImplementation(srv usecase.SocialUsecases) *Implementation {
	return &Implementation{
		srv: srv,
	}
}
