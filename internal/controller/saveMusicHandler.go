package controller

import (
	"SoundLink/internal/app/service"
	"SoundLink/pkg/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GenerateSignedURLHandler godoc
// @Summary Генерация подписанных URL-адресов для загрузки и просмотра
// @Description Сгенерируйте предварительно подписанные URL-адреса для загрузки и просмотра музыкальных файлов. А также сохраняет ссылку для просмотра в БД
// @Tags music
// @Produce json
// @Param filename query string true "File name (e.g. audio.mp3)"
// @Param content_type query string false "MIME type (default: audio/mpeg)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/music/add [post]
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
	uploadURL, err := service.GenerateSignedURL(filename, "PUT", contentType)
	if err != nil {
		log.Printf("Signed URL (PUT) error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate upload URL"})
		return
	}

	// URL для просмотра (GET)
	viewURL, err := service.GenerateSignedURL(filename, "GET", contentType)
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

// DeleteMusic godoc
// @Summary Delete a music track
// @Description Delete a music track by its ID
// @Tags music
// @Produce json
// @Param musicId query int true "Music ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/music/delete [delete]
func DeleteMusic(c *gin.Context) {
	musicId := c.Query("musicId")
	_, err := db.DB.Query("DELETE FROM music WHERE id=?", musicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "Music was deleted",
	})
}
