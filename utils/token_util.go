package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(userId *string) (string, error) {
	fmt.Println(*userId)
	k, kError := getSecretKey("AUTH_TOKEN_SECRET")
	if kError != nil {
		return "", kError
	}
	s := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		UserId: *userId,
	})
	token, err := s.SignedString([]byte(k))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken() bool {
	return false
}

func GenerateAuthToken() string {
	return ""
}

func VerifyAuthToken(token *string) (*TokenClaims, error) {
	k, kError := getSecretKey("AUTH_TOKEN_SECRET")
	if kError != nil {
		fmt.Println(kError)
		return nil, kError
	}
	parsedToken, parseError := jwt.ParseWithClaims(*token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(k), nil
	})
	claims, isClaimValid := parsedToken.Claims.(*TokenClaims)
	if !isClaimValid {
		fmt.Println("Invalid claim")
		return nil, errors.New("invalid claim")
	}
	if parseError != nil {
		fmt.Println(parseError)
		return nil, parseError
	}
	return claims, nil
}

func GenerateEmailVerificationToken() string {
	return ""
}

func VerifyEmailVerificationToken() bool {
	return false
}

func GeneratePasswordResetToken() string {
	return ""
}

func VerifyPasswordResetToken() bool {
	return false
}

func getSecretKey(name string) (string, error) {
	k, kExists := os.LookupEnv(name)
	if !kExists {
		message := fmt.Sprintf("%s not found", name)
		return "", errors.New(message)
	}
	return k, nil
}
