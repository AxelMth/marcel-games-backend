package main

import (
	"fmt"
	"log"
	"marcel-games-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.SetTrustedProxies(nil)

    r.POST("/launch", handlers.LaunchHandler)
    r.POST("/end-level", handlers.EndLevelHandler)
    r.GET("/mapbox-token")

    fmt.Println("Starting server at port 8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
