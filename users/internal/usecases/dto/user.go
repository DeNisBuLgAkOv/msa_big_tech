package dto

type User struct {
	ID       string
	Nickname string
	Bio      string
	Avatar   string
}

type SearchByNicknameRequest struct {
	Query string
	Limit int64
}
