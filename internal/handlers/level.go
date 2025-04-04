package handlers

import (
	"context"
	"fmt"
	"marcel-games-backend/internal/repositories"
	"marcel-games-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LevelInfo struct {
	UserID    string `json:"userId"`
	Attempts  int    `json:"attempts"`
	TimeSpent int    `json:"timeSpent"`
}

func LevelHandler(c *gin.Context) {
	var req LevelInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx := context.Background()

	level := repositories.GetLastLevelFromHistory(ctx, req.UserID)

	rank := utils.CalculateRankForLevel(level, req.Attempts, req.TimeSpent)

	_, err := repositories.CreateOneLevelHistory(
		ctx,
		req.UserID,
		level+1,
		req.Attempts,
		req.TimeSpent,
		rank,
	)

	if err != nil {
		fmt.Println("Failed to create level history", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create level history"})
		return
	}

	response := gin.H{"rank": rank, "nextLevel": level + 2, "nextCountryCodes": utils.GetNextLevelCountryCodes(level + 2)}
	c.JSON(http.StatusOK, response)
}
