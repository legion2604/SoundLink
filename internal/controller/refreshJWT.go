package controller

import (
	"SoundLink/internal/app/models"
	"SoundLink/pkg/db"
	"SoundLink/pkg/utils"
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"
)

var refreshSecret = []byte("secret_refresh")
var accessSecret = []byte("secret_access")

func RefreshTokenHandler(c *gin.Context) {
	req := models.RefreshToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	refreshToken, _ := c.Cookie("refresh_token")
	userPass, err := db.RefreshToken(req.Email, refreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	newAccessToken, err := utils.CreateAccessToken(userPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}
