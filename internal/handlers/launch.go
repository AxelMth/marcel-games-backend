package handlers

import (
	"context"
	"fmt"
	"marcel-games-backend/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func LaunchHandler(c *gin.Context) {
	var req LaunchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx := context.Background()

	user, err := repositories.UpsertOneUser(ctx, req.DeviceUUID)

	if err != nil {
		fmt.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	_, err = repositories.UpsertOneUserDevice(
		ctx,
		user.ID,
		req.Brand,
		req.DeviceType,
		req.IsDevice,
		req.Manufacturer,
		req.ModelName,
		req.OsName,
		req.OsVersion,
	)

	if err != nil {
		fmt.Println("Failed to create user device", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user device"})
		return
	}

	currentLevel := repositories.GetLastLevelFromHistory(ctx, user.ID)

	response := map[string]interface{}{"userId": user.ID, "level": currentLevel + 1}
	c.JSON(http.StatusOK, response)
}
