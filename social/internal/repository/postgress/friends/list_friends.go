package repos_friends

import "context"

func (s *FriendsRepository) ListFriends(ctx context.Context, userID string, limit int, cursor string) ([]string, string, error) {

	return []string{"asdasdas", "sdasdasdasdas"}, "sdasdasdasdas", nil
}
