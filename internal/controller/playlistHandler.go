package controller

import (
	"SoundLink/internal/app/models"
	"SoundLink/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreatePlaylist godoc
// @Summary Создать плейлист
// @Description Создаёт новый плейлист для текущего пользователя
// @Tags playlist
// @Accept json
// @Produce json
// @Param request body models.Playlist true "Данные плейлиста"
// @Success 200 {object} map[string]interface{} "Плейлист успешно создан"
// @Failure 400 {object} map[string]string "Ошибка при парсинге запроса"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Security AccessToken
// @Router /api/playlist/add [post]
func CreatePlaylist(c *gin.Context) {
	var req models.Playlist

	userId := c.GetInt("userId")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := db.DB.Exec(
		"INSERT INTO playlist (playlist_name, creator_id, category) VALUES (?, ?, ?)",
		req.PlaylistName, userId, req.Category,
	)
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"massage":     "Playlist was saved",
			"playlist_id": id,
		})
	}
}

// GetPlaylistToUserId godoc
// @Summary Получить плейлисты пользователя
// @Description Возвращает список плейлистов по userId
// @Tags playlist
// @Produce json
// @Param UserId query int true "ID пользователя"
// @Success 200 {array} models.PlaylistForm "Список плейлистов"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Security AccessToken
// @Router /api/playlist/get [get]
func GetPlaylistToUserId(c *gin.Context) {
	var userId = c.Query("UserId")
	rows, err := db.DB.Query("SELECT id, playlist_name, category FROM playlist WHERE creator_id=?", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var playlists []models.PlaylistForm

	for rows.Next() {
		var pl models.PlaylistForm
		if err := rows.Scan(&pl.Id, &pl.PlaylistName, &pl.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		playlists = append(playlists, pl)
	}

	c.JSON(http.StatusOK, playlists)
}

// DeletePlaylist godoc
// @Summary Удалить плейлист
// @Description Удаляет плейлист по playlistId
// @Tags playlist
// @Produce json
// @Param playlistId query int true "ID плейлиста"
// @Success 200 {object} map[string]string "Плейлист успешно удалён"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Security AccessToken
// @Router /api/playlist [delete]
func DeletePlaylist(c *gin.Context) {
	playlistId := c.Query("playlistId")
	_, err := db.DB.Query("DELETE FROM playlist WHERE id=?", playlistId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "Playlist was deleted",
	})
}
