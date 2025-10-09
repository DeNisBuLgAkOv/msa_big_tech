package usecase

import (
	"context"
	"fmt"
	"msa_big_tech/social/internal/usecase/dto"
)

func (s *UseCase) ListFriends(ctx context.Context, req dto.ListFriendsDto) (*dto.ListFriendsResponse, error) {

	res, cursor, err := s.repo_friends.ListFriends(ctx, req.UsersID, int(req.Limit), *req.Cursor)

	if err != nil {
		return nil, fmt.Errorf("ListFriends ListFriends error: %w", err)
	}

	return &dto.ListFriendsResponse{
		Friends:    res,
		NextCursor: &cursor,
	}, nil
}
