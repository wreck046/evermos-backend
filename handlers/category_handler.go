package handlers

import (
	"github.com/gin-gonic/gin"
)
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

var categories = []Category{}
var categoryID = 1

func CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category := Category{
		ID:   categoryID,
		Name: req.Name,
	}
	categories = append(categories, category)
	categoryID++
	c.JSON(201, category)
}