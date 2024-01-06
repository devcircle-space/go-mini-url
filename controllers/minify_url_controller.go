package controllers

import (
	"fmt"
	"net/http"

	"devcircle.space/mini-url/db"
	"devcircle.space/mini-url/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	idRequired   = "ID is required"
	genericError = "Something went wrong"
	idParseError = "ID is either invalid or cannot be understood"
)

func CreateMinifiedUrl(c *gin.Context) {
	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var u models.UrlPayload
	bindingError := c.BindJSON(&u)
	if bindingError != nil {
		fmt.Println(bindingError)
		c.JSON(http.StatusBadRequest, gin.H{"error": genericError})
		return
	}
	db := db.InitDB()
	u.UserId = userId
	error := u.Create(db)
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, gin.H{"error": genericError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Url minified successfully"})
}

func GetFromMinifiedUrl(c *gin.Context) {
	urlId := c.Params.ByName("id")
	if urlId == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": idRequired})
		return
	}
	parsedId, parseError := primitive.ObjectIDFromHex(urlId)
	if parseError != nil {
		fmt.Println(parseError)
		c.JSON(http.StatusBadRequest, gin.H{"error": idParseError})
		return
	}
	var u models.UrlPayload
	db := db.InitDB()
	u.Id = parsedId
	error := u.Get(db)
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, gin.H{"error": genericError})
		return
	}
	c.Redirect(http.StatusMovedPermanently, u.Link)
}

func UpdateMinifiedUrl(c *gin.Context) {
	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	urlId := c.Params.ByName("id")
	if urlId == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": idRequired})
		return
	}
	parsedId, parseError := primitive.ObjectIDFromHex(urlId)
	if parseError != nil {
		fmt.Println(parseError)
		c.JSON(http.StatusBadRequest, gin.H{"error": idParseError})
		return
	}
	var u models.UrlPayload
	db := db.InitDB()
	u.UserId = userId
	u.Update(&parsedId, db)
}

func DeleteMinifiedUrl(c *gin.Context) {
	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	urlId := c.Params.ByName("id")
	if urlId == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": idRequired})
		return
	}
	parsedId, parseError := primitive.ObjectIDFromHex(urlId)
	if parseError != nil {
		fmt.Println(parseError)
		c.JSON(http.StatusBadRequest, gin.H{"error": idParseError})
		return
	}
	var u models.UrlPayload
	db := db.InitDB()
	u.UserId = userId
	u.Id = parsedId
	error := u.Delete(db)
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, gin.H{"error": genericError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Url deleted successfully"})
}

func GetAllMinifiedUrls(c *gin.Context) {
	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var u models.UrlPayload
	db := db.InitDB()
	u.UserId = userId
	urls, error := u.GetAll(db)
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, gin.H{"error": genericError})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "max-age=3600, public, must-revalidate")
	c.JSON(http.StatusOK, gin.H{"data": urls})
}
