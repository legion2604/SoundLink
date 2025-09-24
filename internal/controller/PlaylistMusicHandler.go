package controller

import (
	"SoundLink/internal/app/models"
	"SoundLink/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddMusicToPlaylist godoc
// @Summary Добавить музыку в плейлист
// @Description Добавить музыку в определенный плейлист
// @Tags playlist
// @Accept json
// @Produce json
// @Param request body models.AddMusicToPlaylist true "Playlist and Music IDs"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/playlist/music [post]
func AddMusicToPlaylist(c *gin.Context) {
	var req models.AddMusicToPlaylist
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	_, err := db.DB.Exec("INSERT INTO playlist_music (playlist_id,music_id) VALUES (?,?)", req.PlaylistID, req.MusicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "music added to playlist",
		})
	}

}

// GetMusicByPlaylist godoc
// @Summary Получить музыку из плейлиста
// @Description Получить все музыки из плейлиста по ID плейлиста
// @Tags playlist
// @Produce json
// @Param playlistID query int true "Playlist ID"
// @Success 200 {array} models.Music
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/playlist/music [get]
func GetMusicByPlaylist(c *gin.Context) {
	playlistID := c.Query("playlistID") // например ?playlistID=1

	rows, err := db.DB.Query(`
        SELECT m.id, m.url
        FROM music m
        JOIN playlist_music pm ON pm.music_id = m.id
        WHERE pm.playlist_id = ?`, playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var musics []models.Music
	for rows.Next() {
		var m models.Music
		if err := rows.Scan(&m.ID, &m.URL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		musics = append(musics, m)
	}

	c.JSON(http.StatusOK, musics)
}

// DeleteMusicInPlaylist godoc
// @Summary Удалить музыку из плейлиста
// @Description Удалить музыкальный трек из определенного плейлиста
// @Tags playlist
// @Accept json
// @Produce json
// @Param request body models.MusicIdPlaylistId true "Playlist and Music IDs"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/playlist/music [delete]
func DeleteMusicInPlaylist(c *gin.Context) {
	var req models.MusicIdPlaylistId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	_, err := db.DB.Query("DELETE FROM playlist_music WHERE playlist_id=? AND music_id=?", req.PlaylistId, req.MusicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "music was deleted",
	})
}
