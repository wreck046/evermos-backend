package handlers

import (
	"evermos-backend/config"
	"evermos-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateTransactionRequest struct {
	AddressID int `json:"address_id"`
	Items []struct {
		ProductID int `json:"product_id"`
		Qty       int `json:"qty"`
	} `json:"items"`
}


func CreateTransaction(c *gin.Context){
	var trreq CreateTransactionRequest
	if err := c.ShouldBindJSON(&trreq); err != nil{
		c.JSON(400, gin.H{
			"error" : "invalid request",
		})
		return
	}
	
	email := c.GetString("email")

	var user models.User
	if err:= config.DB.
	Where("email = ?", email).
	First(&user).
	Error; err != nil{
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	var address models.Address
	if err:= config.DB.
	Where("id = ? AND owner_id = ?", trreq.AddressID, user.ID).
	First(&address).
	Error; err != nil{
		c.JSON(400, gin.H{
			"error": "address not found",
		})
		return
	}
	
	total := 0
	var items []models.TransactionItem
	

	for _, item := range trreq.Items{
		var product models.Product
		if err := config.DB.
		Where("id = ?", item.ProductID).
		First(&product).
		Error; err != nil{
			c.JSON(400, gin.H{
				"error": "product not found",
			})
			return
		}
		subtotal := product.Price * item.Qty
		total += subtotal

		items = append(items, models.TransactionItem{
		ProductID : product.ID,
		ProductName : product.Name,
		Price : product.Price,
		Qty : item.Qty,
		Subtotal : subtotal,
	})
	}

	transaction := models.Transaction{
		UserID: user.ID,
		AddressID: address.ID,
		Total: total,
	}

	if err:= config.DB.
	Create(&transaction).
	Error; err != nil{
		c.JSON(500, gin.H{
			"error": "failed to create transaction",
		})
		return
	}

	for i := range items{
		items[i].TransactionID = transaction.ID
		config.DB.Create(&items[i])
	}
	
	c.JSON(201, gin.H{
		"transaction": transaction,
		"items": items,
	})
}