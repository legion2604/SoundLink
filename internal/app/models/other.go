package models

type UserWithTokens struct {
	IsInDB
	TokenRes
}
type StatusWithTokens struct {
	StatusR
	TokenRes
}
