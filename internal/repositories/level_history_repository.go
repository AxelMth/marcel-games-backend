package repositories

import (
	"context"
	"marcel-games-backend/db"
)

func CreateOneLevelHistory(
	ctx context.Context,
	userID string,
	level int,
	attempts int,
	timeSpent int,
	rank int,
) (*db.LevelHistoryModel, error) {
	levelHistory, err := db.Client().LevelHistory.CreateOne(
		db.LevelHistory.Level.Set(level),
		db.LevelHistory.Attempts.Set(attempts),
		db.LevelHistory.TimeSpent.Set(timeSpent),
		db.LevelHistory.Rank.Set(rank),
		db.LevelHistory.User.Link(db.User.ID.Equals(userID)),
	).Exec(ctx)
	return levelHistory, err
}

func GetLastLevelFromHistory(
	ctx context.Context,
	userID string,
) int {
	levelHistory, err := db.Client().LevelHistory.FindFirst(
		db.LevelHistory.UserID.Equals(userID),
	).OrderBy(
		db.LevelHistory.Level.Order(db.DESC),
	).Exec(ctx)
	if err != nil {
		return 0
	}
	return levelHistory.Level
}
