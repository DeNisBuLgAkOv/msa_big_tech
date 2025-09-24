package transport

import (
	"context"
	users "msa_big_tech/users/pkg/proto/v1"
)

type Implementation struct {
	users.UnimplementedUsersServiceServer
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

func (s *Implementation) CreateProfile(ctx context.Context, req *users.CreateProfileRequest) (*users.CreateProfileResponse, error) {
	bio := "Булгаков Ден"
	avatar := "url"
	return &users.CreateProfileResponse{
		UserProfile: &users.UserProfile{
			UserId:    "112e1-123123-123123123213",
			Nickname:  "Den",
			Bio:       &bio,
			AvatarUrl: &avatar,
		},
	}, nil
}

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
