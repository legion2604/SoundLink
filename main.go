package main

import (
	"SoundLink/internal/controller"
	"SoundLink/internal/controller/middleware"
	"SoundLink/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDB()
	r.POST("/login", controller.LoginHandler)
	r.GET("/profile", controller.ProfileHandler)
	r.POST("/refresh", controller.RefreshTokenHandler)
	// Защищённые роуты
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{

	}

	r.Run(":8080")
}
