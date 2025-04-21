package main

import (
	"fmt"
	"log"
	"marcel-games-backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	r.SetTrustedProxies(nil)

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
