package handlers

import (
	"github.com/gin-gonic/gin"

	"evermos-backend/config"
	"evermos-backend/models"
	"evermos-backend/utils"
)

//mapping saat belum menggunakan database. uji coba logic
var users = map[string]string{}

type RegisterRequest struct {
	Email	string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct{
	Email	string `json:"email"`
	Password string `json:"password"`	
}

type Store struct {
	ID	int `json:"id"`
	Name string `json:"name"`
	Owner string `json:"owner"`
}

var stores = map[string]Store{}
var storeID = 1

func Register(c *gin.Context){
	var req RegisterRequest
	var existingUser models.User

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(400, gin.H{
			"error" : "invalid request",
		})
		return
	}

	err := config.DB.
	Where("email = ?", req.Email).
	First(&existingUser).
	Error

	if err == nil{
		c.JSON(400, gin.H{
			"error" : "email is already used",
		})
		return
	}

	user := models.User{
		Email: req.Email,
		Password: req.Password,
		IsAdmin: false,
	}

	config.DB.Create(&user)

	store := models.Store{
		Name : "Toko " + user.Email ,
		OwnerID :  user.ID,
	}
	config.DB.Create(&store)
	
	c.JSON(201, gin.H{
		"message": "user registered successfully",
	})	
}

func Login(c *gin.Context){
	var login LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil{
		c.JSON(400, gin.H{
			"error" : "invalid request",
		})
		return
	}

	var user models.User
	err := config.DB.
	Where("email = ?", login.Email).
	First(&user).
	Error

	if err != nil {
		c.JSON(401, gin.H{
			"error" : "email not registered",
		})
		return
	}

	if user.Password != login.Password {
		c.JSON(401, gin.H{
			"error" : "wrong password",
		})
		return
	}

	token, err := utils.GenerateToken(login.Email)
	if err != nil{
		c.JSON(500, gin.H{
			"error" : "failed to generate token",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "user logged in successfully",
		"token" : token,
	})


}