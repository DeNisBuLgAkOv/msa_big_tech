package users_grpc

import (
	"context"
	users "msa_big_tech/users/pkg/proto/v1"
)

func (s *Implementation) GetProfileByNickname(ctx context.Context, req *users.GetProfileByNicknameRequest) (*users.GetProfileByNicknameResponse, error) {
	bio := "Булгаков Ден"
	avatar := "url"
	return &users.GetProfileByNicknameResponse{
		UserProfile: &users.UserProfile{
			UserId:    "112e1-123123-123123123213",
			Nickname:  "Den",
			Bio:       &bio,
			AvatarUrl: &avatar,
		},
	}, nil
}
