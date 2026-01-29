package handlers

import (
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID        int    `json:"id"`
	Owner     string `json:"owner"`
	AddressID int    `json:"address_id"`
	Total     int    `json:"total"`
}

type TransactionItem struct {
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	Price         int    `json:"price"`
	Qty           int    `json:"qty"`
	Subtotal      int    `json:"subtotal"`
}

var transactions = []Transaction{}
var transactionItems = []TransactionItem{}
var transactionID = 1

type CreateTransactionRequest struct {
	AddressID int `json:"address_id"`
	Items []struct {
		ProductID int `json:"product_id"`
		Qty       int `json:"qty"`
	} `json:"items"`
}


func CreateTransaction(c *gin.Context){
	trreq := CreateTransactionRequest{}
	c.ShouldBindJSON(&trreq)
	email := c.GetString("email")

	addressFound := false

	for _, addr := range addresses{
		if addr.ID == trreq.AddressID && addr.Owner == email{
			addressFound = true
			break
		}
	}

	if !addressFound {
    c.JSON(400, gin.H{
        "error": "address not found",
    })
    return
	}
	total := 0
	var items = []TransactionItem{}

	for _, item := range trreq.Items{
		productFound := (*Product)(nil)
		for _, p := range products{
			if p.ID == item.ProductID{
				productFound = &p
				break
			}
		}

		if productFound == nil {
    		c.JSON(400, gin.H{
        	"error": "product not found",
    		})
    	return
		}

		subtotal := productFound.Price * item.Qty
		total += subtotal

		transactionItem := TransactionItem{
		ProductID : productFound.ID,
		ProductName : productFound.Name,
		Price : productFound.Price,
		Qty : item.Qty,
		Subtotal : subtotal,
		}

		items = append(items, transactionItem)
	}

	transaction := Transaction{
		ID : transactionID,
		Owner : email,
		AddressID : trreq.AddressID,
		Total : total,
	}

	transactions = append(transactions, transaction)

	for i := range items{
		items[i].TransactionID = transactionID
		transactionItems = append(transactionItems, items[i])
	}
	transactionID++

	c.JSON(201, gin.H{
		"transaction": transaction,
		"items": items,
	})
}