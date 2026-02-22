package main

import (
	"fmt"
	"log"
	"marcel-games-backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// CORS: allow localhost and production frontend; required for browser preflight OPTIONS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://earthunt.com", "https://www.earthunt.com"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/launch", handlers.LaunchHandler)

	r.GET("/level", handlers.GetLevelHandler)
	// deprecated
	r.POST("/end-level", handlers.FinishLevelHandler)
	r.POST("/level", handlers.FinishLevelHandler)

	fmt.Println("Starting server at port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
