package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/timecrunch101/goirc/internal/services/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.Connect()

	fmt.Println("Goodbye, World!")

}
