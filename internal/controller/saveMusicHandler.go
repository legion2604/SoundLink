package controller

import (
	"SoundLink/pkg/db"
	"cloud.google.com/go/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

// GenerateSignedURLHandler генерирует URL для загрузки (PUT) и просмотра (GET).
func GenerateSignedURLHandler(c *gin.Context) {
	// ... (код для userId и filename)
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filename is required"})
		return
	}

	// Для примера, получим тип контента из запроса
	// Например: /generate-signed-url?filename=audio.mp3&content_type=audio/mpeg
	contentType := c.Query("content_type")
	if contentType == "" {
		// Установите тип по умолчанию, если он не был предоставлен
		contentType = "audio/mpeg"
	}

	filename = fmt.Sprintf("%d_%s", userId, filename)

	// URL для загрузки (PUT)
	uploadURL, err := generateSignedURL(filename, "PUT", contentType)
	if err != nil {
		log.Printf("Signed URL (PUT) error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate upload URL"})
		return
	}

	// URL для просмотра (GET)
	viewURL, err := generateSignedURL(filename, "GET", contentType)
	if err != nil {
		log.Printf("Signed URL (GET) error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate view URL"})
		return
	}
	result, _ := db.DB.Exec("INSERT INTO music (url, creatorID) VALUES (?, ?)", viewURL, userId)
	if err != nil {
		log.Println("DB insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file info"})
		return
	}
	id, _ := result.LastInsertId()

	c.JSON(http.StatusOK, gin.H{
		"upload_url": uploadURL,
		"view_url":   viewURL,
		"music_id":   id,
	})
}

// generateSignedURL создает подписанный URL для указанного HTTP-метода.
func generateSignedURL(objectName, method, contentType string) (string, error) {
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("BUCKET_NAME is empty")
	}

	credsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credsFile == "" {
		return "", fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS is empty")
	}

	jsonKey, err := os.ReadFile(credsFile)
	if err != nil {
		return "", err
	}

	var sa struct {
		ClientEmail string `json:"client_email"`
		PrivateKey  string `json:"private_key"`
	}
	if err := json.Unmarshal(jsonKey, &sa); err != nil {
		return "", err
	}

	opts := &storage.SignedURLOptions{
		GoogleAccessID: sa.ClientEmail,
		PrivateKey:     []byte(sa.PrivateKey),
		Method:         method,
		Expires:        time.Now().Add(15 * time.Minute),
	}

	// Для PUT-запросов Content-Type обязателен
	if method == "PUT" {
		opts.ContentType = contentType
	}

	// Здесь самое важное изменение для просмотра
	if method == "GET" {
		// Указываем, что ответ должен иметь заголовок Content-Type
		// и Content-Disposition: inline, чтобы браузер отобразил файл
		opts.Headers = []string{
			"Content-Type:" + contentType,
			"Content-Disposition:inline",
		}
	}

	url, err := storage.SignedURL(bucketName, objectName, opts)
	if err != nil {
		return "", err
	}

	return url, nil
}
