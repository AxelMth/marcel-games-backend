package handlers

import (
	"context"
	"fmt"
	"marcel-games-backend/internal/repositories"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func MapboxAccessTokenHandler(c *gin.Context) {
	userId := c.Query("userId")

    ctx := context.Background()

    var _, err = repositories.GetUserByID(ctx, userId)

    if err != nil {
        fmt.Println("Failed to get user", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    fmt.Println("User ", userId, " requested Mapbox access token")
    
    c.JSON(http.StatusOK, gin.H{"accessToken": os.Getenv("MAPBOX_ACCESS_TOKEN")})
}
