package handlers

import (
	"go-rest-api/entities"
	"go-rest-api/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlayerInfoHandler(c *gin.Context) {
	var player entities.PlayerId
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := postgresql.GetPlayerInfo(entities.TableName, player.PlayerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player": response,
	})
}
