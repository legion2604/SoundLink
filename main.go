package main

import (
	"SoundLink/internal/controller"
	"SoundLink/internal/controller/middleware"
	"SoundLink/pkg/db"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // ⚡ origin фронта
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true, // ⚡ обязательно для HttpOnly cookie
		MaxAge:           12 * time.Hour,
	}))

	db.ConnectDB()
	r.POST("/login", controller.LoginHandler)
	r.GET("/profile", controller.ProfileHandler)
	r.POST("/refresh", controller.RefreshTokenHandler)

	// Защищённые роуты
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/upload", controller.UploadFileHandler)

	}
	fmt.Println(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	r.Run(":8080")
}
