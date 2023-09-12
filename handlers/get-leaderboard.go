package handlers

import (
	"go-rest-api/entities"
	"go-rest-api/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeaderboardHandler(c *gin.Context) {
	response, err := postgresql.GetLeaderboard(entities.TableName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"leaderboard": response,
	})
}
