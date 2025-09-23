package controller

import (
	"SoundLink/internal/app/models"
	"SoundLink/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
