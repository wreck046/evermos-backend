package handlers

import (
	"evermos-backend/config"
	"evermos-backend/models"

	"github.com/gin-gonic/gin"
)


type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	category := models.Category{
		Name: req.Name,
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create category"})
		return
	}

	c.JSON(201, category)
}
