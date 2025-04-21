package handlers

import (
	"context"
	"fmt"
	"marcel-games-backend/internal/constants"
	"marcel-games-backend/internal/repositories"
	"marcel-games-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLevelInfo struct {
	UserID    string `form:"userId" binding:"required"`
	GameMode  string `form:"gameMode" binding:"required"`
	Continent string `form:"continent"`
}

type GetLevelInfoResponse struct {
	Level        int      `json:"level"`
	CountryCodes []string `json:"countryCodes"`
}

func GetLevelHandler(c *gin.Context) {
	var req GetLevelInfo
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	ctx := context.Background()

	level := repositories.GetLastLevelFromHistory(ctx, req.UserID)

	currentLevel := level + 1

	var countryCodes []string
	if req.GameMode == "levelOfTheDay" {
		// TODO: Store levelOfTheDay in DB and retrieve it from there
		countryCodes = utils.GetLevelCountryCodesForLevel(currentLevel)
	} else if req.GameMode == "world" {
		countryCodes = utils.GetLevelCountryCodesForLevel(currentLevel)
	} else if req.GameMode == "continents" {
		// TODO: Add continent check
		continent := constants.Continent(req.Continent)
		countryCodes = utils.GetLevelCountryCodesForContinent(currentLevel, continent)
	}

	response := GetLevelInfoResponse{
		Level:        currentLevel,
		CountryCodes: countryCodes,
	}

	c.JSON(http.StatusOK, response)
}

type FinishLevelInfo struct {
	UserID    string `json:"userId"`
	Attempts  int    `json:"attempts"`
	TimeSpent int    `json:"timeSpent"`
}

type FinishLevelResponse struct {
	Rank             int      `json:"rank"`
	NextLevel        int      `json:"nextLevel"`
	NextCountryCodes []string `json:"nextCountryCodes"`
}

func FinishLevelHandler(c *gin.Context) {
	var req FinishLevelInfo
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

	response := FinishLevelResponse{
		Rank:             rank,
		NextLevel:        level + 2,
		NextCountryCodes: utils.GetLevelCountryCodesForLevel(level + 2),
	}

	c.JSON(http.StatusOK, response)
}
