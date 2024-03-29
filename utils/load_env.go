package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("❌ Error loading .env file")
		panic(err)
	}
	fmt.Println("✅ Successfully loaded .env file")
}
