package handlers

import (
	"evermos-backend/config"
	"evermos-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name 	string  `json:"name"`
	Price 	int `json:"price"`
}

func CreateProduct(c *gin.Context){
	var createreqprod CreateProductRequest
	if err := c.ShouldBindJSON(&createreqprod); err != nil{
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	email := c.GetString("email")
	var user models.User
	if err := config.DB.
	Where("email = ?", email).
	First(&user).
	Error; err != nil{
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	var store models.Store
	if err := config.DB.
	Where("owner_id = ?", user.ID).
	First(&store).
	Error; err != nil{
		c.JSON(404, gin.H{
			"error": "store not found",
		})
		return
	}

	product := models.Product{
		Name: createreqprod.Name,
		Price: createreqprod.Price,
		StoreID: store.ID,
	}

	config.DB.Create(&product)

	c.JSON(201, product)
}

func MyProducts(c *gin.Context) {
	email := c.GetString("email")

	var user models.User
	if err := config.DB.
		Where("email = ?", email).
		First(&user).
		Error; err != nil {

		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	var store models.Store
	if err := config.DB.
		Where("owner_id = ?", user.ID).
		First(&store).
		Error; err != nil {

		c.JSON(404, gin.H{
			"error": "store not found",
		})
		return
	}

	var products []models.Product
	config.DB.
		Where("store_id = ?", store.ID).
		Find(&products)

	c.JSON(200, products)
}
