package users_grpc

import (
	"context"
	users "msa_big_tech/users/pkg/proto/v1"
)

func (s *Implementation) SearchByNickname(ctx context.Context, req *users.SearchByNicknameRequest) (*users.SearchByNicknameResponse, error) {
	bio := "Булгаков Ден"
	avatar := "url"
	profiles := []*users.UserProfile{
		{
			UserId:    "112e1-123123-123123123213",
			Nickname:  "Den",
			Bio:       &bio,
			AvatarUrl: &avatar,
		},
	}
	return &users.SearchByNicknameResponse{
		Profiles: profiles,
	}, nil
}
