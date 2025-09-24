package controller

import "github.com/gin-gonic/gin"

// Logout godoc
// @Summary Выход пользователя
// @Description Удаляет access_token и refresh_token из cookies
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string "Сообщение об успешном выходе"
// @Router /api/logout [get]
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
