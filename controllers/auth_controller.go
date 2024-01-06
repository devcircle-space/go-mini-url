package controllers

import (
	"fmt"
	"net/http"

	"devcircle.space/mini-url/db"
	"devcircle.space/mini-url/models"
	"devcircle.space/mini-url/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := db.InitDB()
	userData, err := user.Find(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValidPassword := utils.ComparePasswordHash(user.Password, userData.Password)
	if !isValidPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}
	_id := userData.Id.Hex()
	token, tokenError := utils.CreateToken(&_id)
	if tokenError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tokenError.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": token})
}
func Register(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := db.InitDB()
	if err := user.Create(db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func Logout(c *gin.Context)               {}
func GetResetPasswordLink(c *gin.Context) {}
func ResetPassword(c *gin.Context)        {}
func VerifyUser(c *gin.Context) {
	id := c.GetString("user_id")
	fmt.Println(id)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	db := db.InitDB()
	var user models.UserLogin
	userData, findError := user.Find(db)
	if findError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": findError.Error()})
		return
	}
	if userData.Id.Hex() != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
func DeleteAccount(c *gin.Context) {}
