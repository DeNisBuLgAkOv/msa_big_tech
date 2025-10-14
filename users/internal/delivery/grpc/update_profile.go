package users_grpc

import (
	"context"
	users "msa_big_tech/users/pkg/proto/v1"
)

func (s *Implementation) UpdateProfile(ctx context.Context, req *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {
	bio := "Обновленная биография"
	avatar := "updated_url"
	return &users.UpdateProfileResponse{
		UserProfile: &users.UserProfile{
			UserId:    "112e1-123123-123123123213",
			Nickname:  "Den",
			Bio:       &bio,
			AvatarUrl: &avatar,
		},
	}, nil
}
