package utils

import (
	"context"
	"marcel-games-backend/db"
)

func CalculateRankForLevel(level int, attempts int, timeSpent int) int {
	ctx := context.Background()

	levelHistories, err := db.Client().LevelHistory.FindMany(
		db.LevelHistory.Level.Equals(level),
	).Exec(ctx)
	if err != nil {
		return 1
	}

	currentScore := calculateScore(attempts, timeSpent)

	betterThan := 0
	for _, history := range levelHistories {
		otherScore := calculateScore(history.Attempts, history.TimeSpent)
		if currentScore > otherScore {
			betterThan++
		}
	}

	if len(levelHistories) > 0 {
		return (betterThan * 100) / len(levelHistories)
	}

	return 100
}

func calculateScore(attempts int, timeSpent int) int {
	return 100 - (attempts * 2) - (timeSpent / 10)
}
