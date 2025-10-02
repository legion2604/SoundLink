package main

import (
	"SoundLink/docs"
	_ "SoundLink/docs"
	"SoundLink/internal/router"
	"SoundLink/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден")
	}

	db.ConnectDB()
	r := routes.SetupRouter()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.0.102:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	setSwagger()

	r.Run("0.0.0.0:8080")
}

func setSwagger() {
	docs.SwaggerInfo.Title = "SoundLink API"
	docs.SwaggerInfo.Description = "Документация API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
}
