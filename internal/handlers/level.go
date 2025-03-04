package handlers

import (
	"context"
	"fmt"
	"marcel-games-backend/internal/constants"
	"marcel-games-backend/internal/repositories"
	"marcel-games-backend/pkg/utils"
	"math/rand"
	"net/http"
	"slices"

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

	response := gin.H{"rank": rank, "nextLevel": level + 2, "nextCountryCodes": getNextLevelCountryCodes(level + 2)}
	c.JSON(http.StatusOK, response)
}

func getNextLevelCountryCodes(level int) []string {
	sorted := sortCountriesByArea(constants.Countries)

	countryCount := getNumberOfCountries(level)
	countrySelectWindow := getCountrySelectWindow(level)
	availableCountries := sorted[:countryCount+countrySelectWindow]

	result := make([]string, 0, countryCount)
	indices := rand.Perm(len(availableCountries))[:countryCount]
	for _, idx := range indices {
		result = append(result, availableCountries[idx].Country)
	}

	return result
}

func sortCountriesByArea(countries []constants.Country) []constants.Country {
	sorted := make([]constants.Country, len(constants.Countries))
	copy(sorted, constants.Countries)
	slices.SortFunc(sorted, func(i, j constants.Country) int {
		if i.Area > j.Area {
			return -1
		}
		if i.Area < j.Area {
			return 1
		}
		return 0
	})
	return sorted
}

func getCountrySelectWindow(level int) int {
	switch {
	case level <= 15:
		return 20
	case level <= 30:
		return 30
	case level <= 50:
		return 50
	case level <= 100:
		return 75
	case level <= 250:
		return 100
	case level <= 500:
		return 150
	case level <= 1000:
		return 200
	default:
		return 200
	}
}

func getNumberOfCountries(level int) int {
	switch {
	case level <= 15:
		return 1
	case level <= 30:
		return rand.Intn(3) + 1
	case level <= 50:
		return rand.Intn(3) + 2
	case level <= 100:
		return rand.Intn(4) + 2
	case level <= 250:
		return rand.Intn(6) + 5
	case level <= 500:
		return rand.Intn(8) + 8
	case level <= 1000:
		return rand.Intn(4) + 12
	default:
		return rand.Intn(6) + 15
	}
}
