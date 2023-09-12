package handlers

import (
	"go-rest-api/entities"
	"go-rest-api/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePlayerInfoHandler(c *gin.Context) {
	playerId := c.Param("id")
	var player entities.PlayerInfo
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if player.PlayerName != "" {
		err := postgresql.UpdatePlayerName(entities.TableName, playerId, player.PlayerName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if player.Score != "" {
		err := postgresql.UpdatePlayerScore(entities.TableName, playerId, player.Score)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "player info has been successfully updated",
	})
}
