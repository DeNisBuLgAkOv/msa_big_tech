package repos_friends

import "database/sql"

type FriendsRepository struct {
	db *sql.DB
}

func NewFriendsRepository(db *sql.DB) *FriendsRepository {
	return &FriendsRepository{db: db}
}
