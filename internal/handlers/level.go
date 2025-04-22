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
	UserID string `form:"userId" binding:"required"`
	// TODO: Add game mode validation
	GameMode string `form:"gameMode" binding:"required"`
	// TODO: Add continent validation
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

	level := repositories.GetLastLevelFromHistory(ctx, req.UserID, req.GameMode, req.Continent)

	currentLevel := level + 1

	countryCodes := getCountryCodes(req.GameMode, req.Continent, currentLevel)

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
	// TODO: Add game mode validation
	GameMode string `json:"gameMode"`
	// TODO: Add continent validation
	Continent string `json:"continent"`
	// TODO: Add country codes validation
	CountryCodes []string `json:"countryCodes"`
}

type FinishLevelResponse struct {
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

	level := repositories.GetLastLevelFromHistory(ctx, req.UserID, req.GameMode, req.Continent)

	_, err := repositories.CreateOneLevelHistory(
		ctx,
		req.UserID,
		level+1,
		req.Attempts,
		req.TimeSpent,
		req.GameMode,
		req.Continent,
		req.CountryCodes,
	)

	if err != nil {
		fmt.Println("Failed to create level history", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create level history"})
		return
	}

	nextLevel := level + 2
	countryCodes := getCountryCodes(req.GameMode, req.Continent, nextLevel)

	response := FinishLevelResponse{
		NextLevel:        nextLevel,
		NextCountryCodes: countryCodes,
	}

	c.JSON(http.StatusOK, response)
}

func getCountryCodes(gameMode string, continent string, level int) []string {
	var countryCodes []string
	if gameMode == "LEVEL_OF_THE_DAY" {
		// TODO: Store levelOfTheDay in DB and retrieve it from there
		countryCodes = utils.GetLevelCountryCodesForLevel(level)
	} else if gameMode == "WORLD" {
		countryCodes = utils.GetLevelCountryCodesForLevel(level)
	} else if gameMode == "CONTINENTS" {
		// TODO: Add continent check (e.g. Africa, Americas, Asia, Europe, Oceania)
		continent := constants.Continent(continent)
		countryCodes = utils.GetLevelCountryCodesForContinent(level, continent)
	}
	return countryCodes
}
