package handlers

import (
	"github.com/gin-gonic/gin"
)

type UserProfile struct{
	Email string `json:"email"`
}

func Me(c *gin.Context){
	email := c.GetString("email")
	userProfile := UserProfile{
		Email : email,
	}
	c.JSON(200, userProfile)
}