package service

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// generateSignedURL создает подписанный URL для указанного HTTP-метода.

func GenerateSignedURL(objectName, method, contentType string) (string, error) {
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
