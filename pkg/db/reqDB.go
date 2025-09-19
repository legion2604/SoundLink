package db

import (
	"SoundLink/internal/app/models"
	"database/sql"
	"log"
)

func VerificationUser(response models.IsInDB, req models.VerificationRequest) (models.IsInDB, error) {

	var id int
	err := DB.QueryRow("SELECT id FROM users WHERE email = ?", req.Email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			response.IsInData = false
			return response, nil
		} else {
			// General database error
			log.Printf("Database error in handleVerification: %v", err)
			return response, err
		}

	}
	response.IsInData = true
	return response, nil
}

func AddUser(data models.SignupJson) error {
	_, err := DB.Exec("INSERT INTO users (name, surname, email, password) VALUES (?, ?, ?, ?)",
		data.Name, data.Surname, data.Email, data.Password)
	if err != nil {
		log.Println("Signup error:", err)
		return err
	}

	return nil
}

func SaveRefreshToken(refreshToken string, email string) {
	_, err := DB.Exec("UPDATE users SET token = ? WHERE email = ?", refreshToken, email)
	if err != nil {
		log.Println("Signup error:", err)
	}
}

func GetPassword(email string, password string) bool {
	var isValid bool
	err := DB.QueryRow(`
    SELECT CASE WHEN password = ? THEN TRUE ELSE FALSE END 
    FROM users 
    WHERE email = ?;
`, password, email).Scan(&isValid)

	if err != nil {
		log.Println("Error:", err)
	}

	return isValid
}

func RefreshToken(email string, refreshToken string) (string, error) {
	var userPass string
	err := DB.QueryRow(
		"SELECT password FROM users WHERE email = ? AND token = ?",
		email, refreshToken,
	).Scan(&userPass)
	return userPass, err
}
