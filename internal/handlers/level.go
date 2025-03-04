package handlers

import (
	"context"
	"fmt"
	"marcel-games-backend/internal/constants"
	"marcel-games-backend/internal/repositories"
	"marcel-games-backend/pkg/utils"
	"math/rand"
	"net/http"
	"sort"

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
	count := getNumberOfCountries(level)

	sorted := make([]constants.Country, len(constants.Countries))
	copy(sorted, constants.Countries)

	levelFactor := float64(level) / 10.0
	sort.Slice(sorted, func(i, j int) bool {
		weight := rand.Float64() < levelFactor
		if weight {
			return sorted[i].Area < sorted[j].Area
		}
		return sorted[i].Area > sorted[j].Area
	})

	startIdx := rand.Intn(len(sorted) / 3)
	availableCountries := sorted[startIdx:]
	if len(availableCountries) < count {
		availableCountries = sorted
	}

	result := make([]string, 0, count)
	indices := rand.Perm(len(availableCountries))[:count]
	for _, idx := range indices {
		result = append(result, availableCountries[idx].Code)
	}

	return result
}

func getNumberOfCountries(level int) int {
	maxCountries := 20
	count := rand.Intn(maxCountries) + 1
	return count
}
