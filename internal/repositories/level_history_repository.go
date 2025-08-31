package repositories

import (
	"context"
	"marcel-games-backend/db"
	"time"
)

func CreateOneLevelHistory(
	ctx context.Context,
	userID string,
	level int,
	attempts int,
	timeSpent int,
	gameMode string,
	continent string,
	countryCodes []string,
) (*db.LevelHistoryModel, error) {
	levelHistory, err := db.Client().LevelHistory.CreateOne(
		db.LevelHistory.Level.Set(level),
		db.LevelHistory.Attempts.Set(attempts),
		db.LevelHistory.TimeSpent.Set(timeSpent),
		db.LevelHistory.User.Link(db.User.ID.Equals(userID)),
		db.LevelHistory.GameMode.Set(db.GameMode(gameMode)),
		db.LevelHistory.Continent.Set(db.Continent(continent)),
		db.LevelHistory.CountryCodes.Set(countryCodes),
	).Exec(ctx)
	return levelHistory, err
}

func GetLastLevelFromHistory(
	ctx context.Context,
	userID string,
	gameMode string,
	continent string,
) int {
	levelHistory, err := db.Client().LevelHistory.FindFirst(
		db.LevelHistory.UserID.Equals(userID),
		db.LevelHistory.GameMode.Equals(db.GameMode(gameMode)),
		db.LevelHistory.Continent.Equals(db.Continent(continent)),
	).OrderBy(
		db.LevelHistory.Level.Order(db.DESC),
	).Exec(ctx)
	if err != nil {
		return 0
	}
	return levelHistory.Level
}

// GetUserDailyLevelCount returns the number of daily levels completed by a user
func GetUserDailyLevelCount(ctx context.Context, userID string) int {
	count, err := db.Client().LevelHistory.FindMany(
		db.LevelHistory.UserID.Equals(userID),
		db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
	).Exec(ctx)
	if err != nil {
		return 0
	}
	return len(count)
}

// GetUserRankForLastDailyLevel returns the user's rank for the most recent daily level
func GetUserRankForLastDailyLevel(ctx context.Context, userID string) (int, error) {
	// Get user's last daily level
	userLastLevel, err := db.Client().LevelHistory.FindFirst(
		db.LevelHistory.UserID.Equals(userID),
		db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
	).OrderBy(
		db.LevelHistory.CreatedAt.Order(db.DESC),
	).Exec(ctx)

	if err != nil {
		return 0, err
	}

	// Get the date range for that level
	levelDate := userLastLevel.CreatedAt
	dayStart := time.Date(levelDate.Year(), levelDate.Month(), levelDate.Day(), 0, 0, 0, 0, levelDate.Location())
	dayEnd := dayStart.Add(24 * time.Hour)

	// Count users with better performance (fewer attempts or same attempts but less time)
	// Since Prisma Go doesn't support OR directly, we need to use two separate queries
	// First, get users with fewer attempts
	fewerAttempts, err := db.Client().LevelHistory.FindMany(
		db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
		db.LevelHistory.CreatedAt.Gte(dayStart),
		db.LevelHistory.CreatedAt.Lt(dayEnd),
		db.LevelHistory.Attempts.Lt(userLastLevel.Attempts),
	).Exec(ctx)

	if err != nil {
		return 0, err
	}

	// Then, get users with same attempts but less time
	sameAttemptsLessTime, err := db.Client().LevelHistory.FindMany(
		db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
		db.LevelHistory.CreatedAt.Gte(dayStart),
		db.LevelHistory.CreatedAt.Lt(dayEnd),
		db.LevelHistory.Attempts.Equals(userLastLevel.Attempts),
		db.LevelHistory.TimeSpent.Lt(userLastLevel.TimeSpent),
	).Exec(ctx)

	if err != nil {
		return 0, err
	}

	// Combine results and remove duplicates
	betterUsersMap := make(map[string]bool)
	for _, history := range fewerAttempts {
		betterUsersMap[history.UserID] = true
	}
	for _, history := range sameAttemptsLessTime {
		betterUsersMap[history.UserID] = true
	}

	// Rank is number of unique better users + 1
	return len(betterUsersMap) + 1, nil
}

// DailyLevelStats holds statistics for a daily level
type DailyLevelStats struct {
	Date      time.Time
	Rank      int
	Attempts  int
	TimeSpent int
}

// GetUserGlobalDailyRank calculates the user's global rank based on all daily levels
func GetUserGlobalDailyRank(ctx context.Context, userID string) (int, error) {
	// Get all user's daily levels
	userLevels, err := db.Client().LevelHistory.FindMany(
		db.LevelHistory.UserID.Equals(userID),
		db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
	).OrderBy(
		db.LevelHistory.CreatedAt.Order(db.ASC),
	).Exec(ctx)

	if err != nil || len(userLevels) == 0 {
		return 0, err
	}

	// Calculate average rank across all daily levels
	totalScore := 0.0
	for _, level := range userLevels {
		// Get the date range for this level
		levelDate := level.CreatedAt
		dayStart := time.Date(levelDate.Year(), levelDate.Month(), levelDate.Day(), 0, 0, 0, 0, levelDate.Location())
		dayEnd := dayStart.Add(24 * time.Hour)

		// Count users with better performance for this specific day
		// Split into two queries
		fewerAttempts, _ := db.Client().LevelHistory.FindMany(
			db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
			db.LevelHistory.CreatedAt.Gte(dayStart),
			db.LevelHistory.CreatedAt.Lt(dayEnd),
			db.LevelHistory.Attempts.Lt(level.Attempts),
		).Exec(ctx)

		sameAttemptsLessTime, _ := db.Client().LevelHistory.FindMany(
			db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
			db.LevelHistory.CreatedAt.Gte(dayStart),
			db.LevelHistory.CreatedAt.Lt(dayEnd),
			db.LevelHistory.Attempts.Equals(level.Attempts),
			db.LevelHistory.TimeSpent.Lt(level.TimeSpent),
		).Exec(ctx)

		// Combine and deduplicate
		betterUsersMap := make(map[string]bool)
		for _, history := range fewerAttempts {
			betterUsersMap[history.UserID] = true
		}
		for _, history := range sameAttemptsLessTime {
			betterUsersMap[history.UserID] = true
		}

		dayRank := len(betterUsersMap) + 1
		// Weight recent performances more heavily
		weight := 1.0
		totalScore += float64(dayRank) * weight
	}

	avgRank := totalScore / float64(len(userLevels))

	// Now compare with other users' average ranks
	allUsers, _ := db.Client().User.FindMany().Exec(ctx)

	betterUsersCount := 0
	for _, otherUser := range allUsers {
		if otherUser.ID == userID {
			continue
		}

		otherUserLevels, _ := db.Client().LevelHistory.FindMany(
			db.LevelHistory.UserID.Equals(otherUser.ID),
			db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
		).Exec(ctx)

		if len(otherUserLevels) == 0 {
			continue
		}

		// Calculate other user's average rank
		otherTotalScore := 0.0
		for _, level := range otherUserLevels {
			levelDate := level.CreatedAt
			dayStart := time.Date(levelDate.Year(), levelDate.Month(), levelDate.Day(), 0, 0, 0, 0, levelDate.Location())
			dayEnd := dayStart.Add(24 * time.Hour)

			// Split into two queries
			fewerAttempts, _ := db.Client().LevelHistory.FindMany(
				db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
				db.LevelHistory.CreatedAt.Gte(dayStart),
				db.LevelHistory.CreatedAt.Lt(dayEnd),
				db.LevelHistory.Attempts.Lt(level.Attempts),
			).Exec(ctx)

			sameAttemptsLessTime, _ := db.Client().LevelHistory.FindMany(
				db.LevelHistory.GameMode.Equals(db.GameMode("LEVEL_OF_THE_DAY")),
				db.LevelHistory.CreatedAt.Gte(dayStart),
				db.LevelHistory.CreatedAt.Lt(dayEnd),
				db.LevelHistory.Attempts.Equals(level.Attempts),
				db.LevelHistory.TimeSpent.Lt(level.TimeSpent),
			).Exec(ctx)

			// Combine and deduplicate
			betterUsersMap := make(map[string]bool)
			for _, history := range fewerAttempts {
				betterUsersMap[history.UserID] = true
			}
			for _, history := range sameAttemptsLessTime {
				betterUsersMap[history.UserID] = true
			}

			dayRank := len(betterUsersMap) + 1
			otherTotalScore += float64(dayRank)
		}

		otherAvgRank := otherTotalScore / float64(len(otherUserLevels))

		// Consider both average rank and number of levels completed
		// Users with more levels completed and better average rank are ranked higher
		if len(otherUserLevels) > len(userLevels) ||
			(len(otherUserLevels) == len(userLevels) && otherAvgRank < avgRank) {
			betterUsersCount++
		}
	}

	return betterUsersCount + 1, nil
}
