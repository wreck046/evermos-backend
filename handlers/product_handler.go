package handlers

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID		int     `json:"id"`
	Name 	string  `json:"name"`
	Price 	int `json:"price"`
	StoreID int  `json:"store_id"`
}

type CreateProductRequest struct {
	Name 	string  `json:"name"`
	Price 	int `json:"price"`
}

var products = []Product{}
var productID = 1

func CreateProduct(c *gin.Context){
	var createreqprod CreateProductRequest
	email := c.GetString("email")

	err := c.ShouldBindJSON(&createreqprod)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	store, exists := stores[email]
	if !exists{
		c.JSON(404, gin.H{
			"error": "store not found/not created",
		})
		return
	}

	product := Product{
		ID : productID,
		Name : createreqprod.Name,
		Price : createreqprod.Price,
		StoreID : store.ID,
	}

	products = append(products, product)
	productID++

	c.JSON(201, product)
}

func MyProducts(c *gin.Context){
	email := c.GetString("email")
	store, exists := stores[email]
	if !exists{
		c.JSON(404, gin.H{
			"error": "store not found/not created",
		})
		return
	}
	var myProducts []Product
	for _, prod := range products{
		if prod.StoreID == store.ID{
			myProducts = append(myProducts, prod)
		}
	}
	c.JSON(200, myProducts)
}