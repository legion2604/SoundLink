package controller

import (
	"SoundLink/internal/app/models"
	"SoundLink/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
