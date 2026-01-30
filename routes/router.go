package routes

import (
	"evermos-backend/handlers"
	"evermos-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	router.GET("/health", middleware.JWTAuth, handlers.HealthCheck)

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.GET("/profile", middleware.JWTAuth, handlers.Me)
	router.GET("/my-store", middleware.JWTAuth, handlers.MyStore)

	router.POST("/products", middleware.JWTAuth, handlers.CreateProduct)
	router.GET("/my-products", middleware.JWTAuth, handlers.MyProducts)

	router.POST("/categories", middleware.JWTAuth, middleware.AdminOnly, handlers.CreateCategory)

	router.POST("/addresses", middleware.JWTAuth, handlers.CreateAddress)
	router.GET("/my-addresses", middleware.JWTAuth, handlers.MyAddresses)

	router.POST("/create-transaction", middleware.JWTAuth, handlers.CreateTransaction)
}
