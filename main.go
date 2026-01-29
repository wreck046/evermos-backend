package main

import (
	"evermos-backend/config"
	"evermos-backend/models"
	"evermos-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&models.User{},
	)
	router.Run()
}