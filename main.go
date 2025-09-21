package main

import (
	"SoundLink/internal/controller"
	"SoundLink/internal/controller/middleware"
	"SoundLink/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	os.Setenv("BUCKET_NAME", "project_sound_link")                                                                    // Замените на имя вашего бакета
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/Users/Win10_Game_OS/IdeaProjects/SoundLink/service-account.json") // Замените на путь к вашему файлу JSON
	db.ConnectDB()
	r.POST("/login", controller.LoginHandler)
	r.GET("/profile", controller.ProfileHandler)
	r.POST("/refresh", controller.RefreshTokenHandler)

	// Защищённые роуты
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/generate-signed-url", controller.GenerateSignedURLHandler)

	}

	r.Run(":8080")
}
