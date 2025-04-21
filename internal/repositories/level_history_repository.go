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
