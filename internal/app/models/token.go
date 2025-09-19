package models

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var RefreshToken struct {
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
}
