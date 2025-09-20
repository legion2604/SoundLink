package controller

import (
	"SoundLink/internal/app/models"
	"SoundLink/internal/app/service"
	"SoundLink/pkg/db"
	_ "SoundLink/pkg/db"
	"SoundLink/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {

	var req models.VerificationRequest

	response := models.IsInDB{IsVer: false, IsInData: false}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.IsVer = service.IsPasswordCorrect(req.Email, req.Password)

	isInDB, userId, _ := db.VerificationUser(response, req) // is user in DB
	if isInDB.IsInData == true && isInDB.IsVer == true {
		refreshToken, _ := utils.CreateRefreshToken(userId)
		accessToken, _ := utils.CreateAccessToken(userId)
		db.SaveRefreshToken(refreshToken, req.Email) // save refresh token in DB to email
		var res = models.UserWithTokens{
			IsInDB:   models.IsInDB{IsVer: true, IsInData: true},
			TokenRes: models.TokenRes{AccessToken: accessToken, RefreshToken: refreshToken},
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, isInDB)
	}
} //success

func ProfileHandler(c *gin.Context) {
	var req models.SignupJson
	res := models.StatusR{Status: false}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, res)
	}
	id, err := db.AddUser(req) // save user data
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		refreshToken, _ := utils.CreateRefreshToken(id)
		accessToken, _ := utils.CreateAccessToken(id)
		db.SaveRefreshToken(refreshToken, req.Email) // save refresh token in DB to email
		var res = models.StatusWithTokens{
			StatusR:  models.StatusR{Status: true},
			TokenRes: models.TokenRes{AccessToken: accessToken, RefreshToken: refreshToken},
		}
		c.JSON(http.StatusOK, res)
	}
} //success
