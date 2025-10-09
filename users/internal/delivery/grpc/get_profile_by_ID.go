package users_grpc

import (
	"context"
	users "msa_big_tech/users/pkg/proto/v1"
)

func (s *Implementation) GetProfileByID(ctx context.Context, req *users.GetProfileByIDRequest) (*users.GetProfileByIDResponse, error) {
	bio := "Булгаков Ден"
	avatar := "url"
	return &users.GetProfileByIDResponse{
		UserProfile: &users.UserProfile{
			UserId:    "112e1-123123-123123123213",
			Nickname:  "Den",
			Bio:       &bio,
			AvatarUrl: &avatar,
		},
	}, nil
}
