package main

import (
	"evermos-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	// router := gin.Default()
	// router.GET("/health", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"status": "healthy",
	// 	})
	// })
	router.Run()
}