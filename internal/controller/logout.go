package controller

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {
	// Устанавливаем пустую куку с истёкшим сроком
	c.SetCookie("access_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}
