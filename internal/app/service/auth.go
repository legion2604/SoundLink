package service

import "SoundLink/pkg/db"

func IsPasswordCorrect(email string, password string) bool {
	return db.GetPassword(email, password)
}
