package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"devcircle.space/mini-url/utils"
	"github.com/gin-gonic/gin"
)

func ValidateAuthToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization missing"})
		return
	}
	token := strings.Split(authHeader, "Bearer ")[1]
	tokenClaims, tokenError := utils.VerifyAuthToken(&token)
	if tokenError != nil {
		fmt.Println("Verification error:", tokenError)
		c.JSON(http.StatusBadRequest, gin.H{"error": tokenError.Error()})
		c.Abort()
		return
	}
	if tokenClaims == nil {
		fmt.Println("Invalid token, no claims")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	c.Set("user_id", tokenClaims.UserId)
	c.Next()
}
