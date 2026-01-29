package handlers

import (
	"github.com/gin-gonic/gin"
)

type Address struct {
	ID		int		`json:"id"`
	Label	string	`json:"label"`
	Detail	string	`json:"detail"`
	City	string	`json:"city"`
	Owner	string	`json:"owner"`
}

type CreateAddressRequest struct {
	Label	string	`json:"label"`
	Detail	string	`json:"detail"`
	City	string	`json:"city"`	
}

var addresses = []Address{}
var addressID = 1

func CreateAddress(c *gin.Context){
	var req CreateAddressRequest
	email := c.GetString("email")
	err := c.ShouldBindJSON(&req)

	if err != nil{
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	address := Address{
		ID		: addressID,
		Label	: req.Label,
		Detail	: req.Detail,
		City	: req.City,
		Owner	: email,
	}
	
	addresses = append(addresses, address)
	addressID++
	c.JSON(201, address)
}

func MyAddresses(c *gin.Context){
	email := c.GetString("email")
	usersAddresses := []Address{}
	for _, addr := range addresses{
		if addr.Owner == email{
			usersAddresses = append(usersAddresses, addr)
		}
	}
	c.JSON(200, usersAddresses)
}