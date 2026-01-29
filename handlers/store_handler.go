package handlers

import (
	"github.com/gin-gonic/gin"
)

func MyStore(c *gin.Context) {
	email := c.GetString("email")
	store, exists := stores[email]
	if !exists{
		c.JSON(404, gin.H{
			"error": "store not found/not created",
		})
		return
	}
	c.JSON(200, store)
}