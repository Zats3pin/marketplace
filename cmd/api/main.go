package main

import (
	"MarketPlaceSolo/internal/config"
	"MarketPlaceSolo/internal/pkg/db"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	cfg := config.Load()
	database := db.Connect(cfg)
	defer database.Close()

	fmt.Printf("postgreSQL connection successful\n")
}
