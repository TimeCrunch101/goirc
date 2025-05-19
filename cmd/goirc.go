package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
	httpserver "github.com/timecrunch101/goirc/internal/services/httpServer"
	"github.com/timecrunch101/goirc/internal/services/irc"
	"github.com/timecrunch101/goirc/internal/services/mysql"
)

var wg sync.WaitGroup

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.Connect()
	wg.Add(1)
	go func() {
		defer wg.Done()
		irc.StartServer()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpserver.Start()

	}()

	wg.Wait()

	fmt.Println("Goodbye, World!")

}
