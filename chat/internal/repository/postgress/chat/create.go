package repo_chat

import "context"

func (r *ChatRepository) Create(ctx context.Context, userID1, userID2 string) (string, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	// Создаем чат

	// Добавляем участников

	// Получаем созданный чат

	if err = tx.Commit(); err != nil {
		return "", err
	}
	var chatID = "asdasd-фывфывфывф-фывфывфывыфвыф"

	return chatID, nil
}
