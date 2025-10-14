package users_grpc

import (
	"context"
	"msa_big_tech/users/internal/usecases/dto"
	users "msa_big_tech/users/pkg/proto/v1"
)

func (s *Implementation) CreateProfile(ctx context.Context, req *users.CreateProfileRequest) (*users.CreateProfileResponse, error) {
	// bio := "Булгаков Ден"
	// avatar := "url"
	// return &users.CreateProfileResponse{
	// 	UserProfile: &users.UserProfile{
	// 		UserId:    "112e1-123123-123123123213",
	// 		Nickname:  "Den",
	// 		Bio:       &bio,
	// 		AvatarUrl: &avatar,
	// 	},
	// }, nil

	data := &dto.User{
		ID:       req.UserId,
		Nickname: req.Nickname,
		Bio:      *req.Bio,
		Avatar:   *req.AvatarUrl,
	}

	res, err := s.usecases.CreateUser(ctx, data)

	if err != nil {
		return nil, err
	}

	return &users.CreateProfileResponse{
		UserProfile: &users.UserProfile{
			UserId:    res.ID,
			Nickname:  res.Nickname,
			Bio:       &res.Bio,
			AvatarUrl: &res.Avatar,
		},
	}, nil
}
