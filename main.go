package main

import (
	"SoundLink/docs"
	_ "SoundLink/docs"
	"SoundLink/internal/controller"
	"SoundLink/internal/controller/middleware"
	"SoundLink/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.POST("/registration", controller.RegistrationHandler)
	r.POST("/refresh", controller.RefreshTokenHandler)

	// Защищённые роуты
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/music/add", controller.GenerateSignedURLHandler)
		auth.POST("/playlist/add", controller.CreatePlaylist)
		auth.POST("/playlist/music/add", controller.AddMusicToPlaylist)
		auth.GET("/playlist/get", controller.GetPlaylistToUserId)
		auth.GET("/playlist/music/get", controller.GetMusicByPlaylist)
		auth.GET("/logout", controller.Logout)
		auth.DELETE("/playlist/delete", controller.DeletePlaylist) // middleware from delete playlist (role)
		auth.DELETE("/playlist/music/delete", controller.DeleteMusicInPlaylist)

	}
	docs.SwaggerInfo.Title = "SoundLink API"
	docs.SwaggerInfo.Description = "Документация API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run("0.0.0.0:8080")
}
