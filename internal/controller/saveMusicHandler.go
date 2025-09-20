package controller

import (
	"SoundLink/pkg/db"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

const bucketName = "project_sound_link" // бакет в GCS

func UploadFileHandler(c *gin.Context) {
	// Получаем userId из контекста (он пришёл из токена)
	userId := c.GetInt("userId") // теперь точно будет int
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Получаем файл из формы
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found"})
		return
	}
	defer file.Close()

	// Создаем имя файла в бакете (можно добавить userId в имя для уникальности)
	filename := fmt.Sprintf("%d_%s", userId, header.Filename)

	// Контекст для GCS
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("GCS client error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GCS client error"})
		return
	}
	defer client.Close()

	// Загружаем файл в бакет
	wc := client.Bucket(bucketName).Object(filename).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Println("Failed to write to GCS:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload"})
		return
	}
	if err := wc.Close(); err != nil {
		log.Println("GCS Close error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Формируем публичную ссылку
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, filename)

	// Сохраняем ссылку в БД
	_, err = db.DB.Exec("INSERT INTO music (url, creatorID) VALUES (?, ?)", publicURL, userId)
	if err != nil {
		log.Println("DB insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file info"})
		return
	}

	// Ответ клиенту
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"url":     publicURL,
		"userId":  userId,
	})
}
