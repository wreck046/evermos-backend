package middleware

import (
	"github.com/gin-gonic/gin"
)
var adminEmail = map[string]bool {
	"admin@email.com": true,
}

func AdminOnly(c *gin.Context) {
	email := c.GetString("email")
	if(!adminEmail[email]){
		c.JSON(403, gin.H{
			"error": "forbidden access",
		})
		c.Abort()
		return
	}
	c.Next()
}