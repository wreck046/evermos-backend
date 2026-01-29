package middleware

import "github.com/gin-gonic/gin"

func AuthDummy(c *gin.Context) {
	header := c.GetHeader("X-User")
	if(header == ""){
		c.JSON(401, gin.H{
			"error": "no user detected",
		})
		c.Abort()
		return
	}
	c.Next()
}