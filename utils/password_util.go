package utils

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) ([]byte, error) {
	cost, doesExists := os.LookupEnv("HASH_COST")
	if !doesExists {
		cost = "10"
	}
	hashCost, err := strconv.Atoi(cost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	hash, hashError := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if hashError != nil {
		fmt.Println(hashError)
		return nil, hashError
	}
	return hash, nil
}

func ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
