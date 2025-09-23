package main

import (
	"SoundLink/internal/controller"
	"SoundLink/internal/controller/middleware"
	"SoundLink/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	db.ConnectDB()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.0.102:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	os.Setenv("BUCKET_NAME", "project_sound_link")                                                                    // Замените на имя вашего бакета
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/Users/Win10_Game_OS/IdeaProjects/SoundLink/service-account.json") // Замените на путь к вашему файлу JSON

	r.POST("/login", controller.LoginHandler)
	r.GET("/profile", controller.ProfileHandler)
	r.POST("/refresh", controller.RefreshTokenHandler)

	// Защищённые роуты
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/createPlaylist", controller.CreatePlaylist)
		auth.GET("/generate-signed-url", controller.GenerateSignedURLHandler)
		auth.POST("/logout", controller.Logout)
		auth.POST("/addMusicToPlaylist", controller.AddMusicToPlaylist)
		auth.GET("/getPlaylistToUserId", controller.GetPlaylistToUserId)
		auth.GET("/getMusicByPlaylist", controller.GetMusicByPlaylist)
		auth.DELETE("/deletePlaylist", controller.DeletePlaylist) // middleware from delete playlist (role)
		auth.DELETE("/deleteMusicInPlaylist", controller.DeleteMusicInPlaylist)

	}
	r.Run("0.0.0.0:8080")
}
