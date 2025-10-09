package repo_request

import "database/sql"

type RequestsRepository struct {
	db *sql.DB
}

func NewRequestsRepository(db *sql.DB) *RequestsRepository {
	return &RequestsRepository{db: db}
}
