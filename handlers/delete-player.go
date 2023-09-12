package handlers

import (
	"go-rest-api/entities"
	"go-rest-api/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePlayerHandler(c *gin.Context) {
	playerId := c.Param("id")

	err := postgresql.DeletePlayer(entities.TableName, playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "player has been deleted",
	})
}
