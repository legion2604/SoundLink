package service

import (
	"SoundLink/pkg/db"
	"log"
)

func SaveFile(userId int, fileURL string) error {
	_, err := db.DB.Exec(
		"INSERT INTO files (user_id, file_url) VALUES (?, ?)",
		userId, fileURL,
	)
	if err != nil {
		log.Println("Save file error:", err)
		return err
	}
	return nil
}
