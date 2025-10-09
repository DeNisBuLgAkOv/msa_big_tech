package repo_chat

import "context"

func (s *ChatRepository) GetUsersByChatID(ctx context.Context, chatID string) ([]string, error) {

	return []string{"sadadas", "sadasdas"}, nil
}
