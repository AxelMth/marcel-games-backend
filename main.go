package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"marcel-games-backend/db"

	"github.com/gin-gonic/gin"
)

var client *db.PrismaClient

func init() {
    var err error
    client = db.NewClient()
    if err = client.Prisma.Connect(); err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}

func main() {
    defer func() {
        if err := client.Prisma.Disconnect(); err != nil {
            log.Fatalf("failed to disconnect from database: %v", err)
        }
    }()

    r := gin.Default()

    r.POST("/launch", launchHandler)
    r.POST("/end-level", endLevelHandler)
    r.GET("/next-level", getNextLevel)

    fmt.Println("Starting server at port 8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}

type LaunchRequest struct {
    DeviceUUID   string `json:"deviceUUID"`
    Brand        string `json:"brand"`
    DeviceType   string `json:"deviceType"`
    IsDevice     bool   `json:"isDevice"`
    Manufacturer string `json:"manufacturer"`
    ModelName    string `json:"modelName"`
    OsName       string `json:"osName"`
    OsVersion    string `json:"osVersion"`
}

func launchHandler(c *gin.Context) {
    var req LaunchRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    ctx := context.Background()

    user, err := client.User.UpsertOne(
        db.User.DeviceUUID.Equals(req.DeviceUUID),
    ).Create(
        db.User.DeviceUUID.Set(req.DeviceUUID),
        db.User.LastLogin.Set(time.Now()),
        db.User.OpenCount.Set(1),
    ).Update(
        db.User.LastLogin.Set(time.Now()),
        db.User.OpenCount.Increment(1),
    ).Exec(ctx)
    if err != nil {
        fmt.Println("Failed to create user", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    _, err = client.UserDevice.UpsertOne(
        db.UserDevice.UserID.Equals(user.ID),
    ).Create(
        db.UserDevice.Brand.Set(req.Brand),
        db.UserDevice.DeviceType.Set(req.DeviceType),
        db.UserDevice.IsDevice.Set(req.IsDevice),
        db.UserDevice.Manufacturer.Set(req.Manufacturer),
        db.UserDevice.ModelName.Set(req.ModelName),
        db.UserDevice.OsName.Set(req.OsName),
        db.UserDevice.OsVersion.Set(req.OsVersion),
        db.UserDevice.User.Link(db.User.ID.Equals(user.ID)),
    ).Update(
        db.UserDevice.Brand.Set(req.Brand),
        db.UserDevice.DeviceType.Set(req.DeviceType),
        db.UserDevice.IsDevice.Set(req.IsDevice),
        db.UserDevice.Manufacturer.Set(req.Manufacturer),
        db.UserDevice.ModelName.Set(req.ModelName),
        db.UserDevice.OsName.Set(req.OsName),
        db.UserDevice.OsVersion.Set(req.OsVersion),
    ).Exec(ctx)
    if err != nil {
        fmt.Println("Failed to create user device", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user device"})
        return
    }

    currentLevel := getLastLevelFromHistory(user.ID)

    response := map[string]interface{}{"userId": user.ID, "level": currentLevel + 1}
    c.JSON(http.StatusOK, response)
}

type LevelInfo struct {
    UserID    string `json:"userId"`
    Attempts  int `json:"attempts"`
    TimeSpent int `json:"timeSpent"`
}

func endLevelHandler(c *gin.Context) {
    var req LevelInfo
    if err := c.ShouldBindJSON(&req); err != nil {
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    ctx := context.Background()

    level := getLastLevelFromHistory(req.UserID)

    _, err := client.LevelHistory.CreateOne(
        db.LevelHistory.Level.Set(level + 1),
        db.LevelHistory.Attempts.Set(req.Attempts),
        db.LevelHistory.TimeSpent.Set(req.TimeSpent),
        db.LevelHistory.Rank.Set(calculateRank(req.Attempts, req.TimeSpent)),
        db.LevelHistory.User.Link(db.User.ID.Equals(req.UserID)),
    ).Exec(ctx)
    if err != nil {
        fmt.Println("Failed to create level history", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create level history"})
        return
    }

    rank := calculateRank(req.Attempts, req.TimeSpent)
    response := map[string]int{"rank": rank}
    c.JSON(http.StatusOK, response)
}

type NextLevel struct {
    UserID    string `json:"userId"`
}

func getNextLevel(c *gin.Context) {
    userId    := c.Query("userId")

    currentLevel := getLastLevelFromHistory(userId)

    response := map[string]int{"nextLevel": currentLevel + 1}
    c.JSON(http.StatusOK, response)
}

func calculateRank(attempts, timeSpent int) int {
    return 100 - (attempts * 2) - (timeSpent / 10)
}

func getLastLevelFromHistory(userID string) (int) {
    ctx := context.Background()
    levelHistory, err := client.LevelHistory.FindFirst(
        db.LevelHistory.UserID.Equals(userID),
    ).OrderBy(
        db.LevelHistory.Level.Order(db.DESC),
    ).Exec(ctx)
    if err != nil {
        return 0
    }
    return levelHistory.Level
}