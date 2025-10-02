package routes

import (
	"SoundLink/internal/controller"
	"SoundLink/internal/controller/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controller.LoginHandler)
	r.POST("/registration", controller.RegistrationHandler)
	r.POST("/refresh", controller.RefreshTokenHandler)

	auth := r.Group("/api", middleware.AuthMiddleware())
	{
		auth.GET("/logout", controller.Logout)

		// Группа для плейлистов
		playlistApi := auth.Group("/playlist")
		{
			playlistApi.POST("/add", controller.CreatePlaylist)
			playlistApi.POST("/music/add", controller.AddMusicToPlaylist)
			playlistApi.GET("/get", controller.GetPlaylistToUserId)
			playlistApi.GET("/music/get", controller.GetMusicByPlaylist)
			playlistApi.DELETE("/delete", controller.DeletePlaylist)
			playlistApi.DELETE("/music/delete", controller.DeleteMusicInPlaylist)
		}

		// Группа для музыки
		musicApi := auth.Group("/music")
		{
			musicApi.POST("/add", controller.GenerateSignedURLHandler)
			musicApi.DELETE("/delete", controller.DeleteMusic)
		}
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
