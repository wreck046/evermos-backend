package main

import (
	"os"

	"evermos-backend/config"
	"evermos-backend/middleware"
	"evermos-backend/models"
	"evermos-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Product{},
		&models.Address{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.Category{},
	)
	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}