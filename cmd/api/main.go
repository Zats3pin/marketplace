package main

import (
	"MarketPlaceSolo/internal/config"
	"MarketPlaceSolo/internal/domain"
	"MarketPlaceSolo/internal/pkg/db"
	"MarketPlaceSolo/internal/service"
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

	//repo := service.NewProductRepository(database)
	//
	//product := &domain.Product{
	//	Name:  "Test product",
	//	Price: 999.99,
	//	Stock: 10,
	//}
	//
	//err := repo.Create(product)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("created product with id:", product.ID)

	repo := service.NewUserRepository(database)

	user := &domain.User{
		Username: "Test product",
		Email:    "user@example.com",
	}

	err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("created product with id:", user.ID)
}
