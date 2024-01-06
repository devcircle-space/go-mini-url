package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
type TokenPayload struct {
	UserId    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
	IssuedAt  int64  `json:"issued_at"`
	IssuedBy  string `json:"issued_by"`
}

func (tp *TokenPayload) CreateToken() (string, error) {
	k, kError := getSecretKey("AUTH_TOKEN_SECRET")
	if kError != nil {
		return "", kError
	}
	s := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"user_id": tp.UserId,
		"exp":     tp.ExpiresAt,
		"iat":     tp.IssuedAt,
		"iss":     tp.IssuedBy,
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
	if parseError != nil {
		fmt.Println(parseError)
		return nil, parseError
	}
	claims, isClaimValid := parsedToken.Claims.(*TokenClaims)
	if !isClaimValid {
		fmt.Println("Invalid claim")
		return nil, jwt.ErrTokenInvalidClaims
	}
	// check if token is expired
	fmt.Println(claims.ExpiresAt.Compare(time.Now()))
	if claims.ExpiresAt.Compare(time.Now()) < 0 {
		fmt.Println("Token expired")
		return nil, jwt.ErrTokenExpired
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
