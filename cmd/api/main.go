package main

import (
	"MarketPlaceSolo/internal/config"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	cfg := config.Load()
	fmt.Printf("server run on port %s\n", cfg.ServerPort)
}
