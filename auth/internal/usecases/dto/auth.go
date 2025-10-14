package usecase_dto

type Register struct {
	Email    string
	Password string
}

type Login struct {
	Email    string
	Password string
}

type AuthRefresh struct {
	ID           string
	RefreshToken string
	AccessToken  string
}
