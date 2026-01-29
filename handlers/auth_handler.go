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

	//dibawah ini merupakan fungsional register user dengan mapping/output JSON untuk testing logic

	// _, exist := users[req.Email]
	// if exist{
	// 	c.JSON(400, gin.H{
	// 		"error" : "email is already used",
	// 	})
	// 	return
	// }

	// users[req.Email] = req.Password
	// namaToko := "Toko " + req.Email
	// stores[req.Email] = Store{
	// 	ID : storeID,
	// 	Name : namaToko,
	// 	Owner : req.Email,
	// }
	// storeID += 1

	c.JSON(200, gin.H{
		"message": "user registered successfully",
	})	
}

func Login(c *gin.Context){
	var login LoginRequest
	err := c.ShouldBindJSON(&login)
	if err != nil{
		c.JSON(400, gin.H{
			"error" : "invalid request",
		})
		return
	}

	storedPassword, exists := users[login.Email]

	if !exists {
		c.JSON(401, gin.H{
			"error" : "email not registered",
		})
		return
	}

	if storedPassword != login.Password {
		c.JSON(401, gin.H{
			"error" : "wrong password",
		})
		return
	}

	token, err := utils.GenerateToken(login.Email)
	
	c.JSON(200, gin.H{
		"message": "user logged in successfully",
		"token" : token,
	})


}