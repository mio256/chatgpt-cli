package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func env() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Can't read: %v\n", err)
	}
}
