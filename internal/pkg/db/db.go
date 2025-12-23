package db

import (
	"database/sql"
	"fmt"
	"log"

	"MarketPlaceSolo/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) *sql.DB {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("failed opening connection to postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed pinging postgres: %v", err)
	}

	return db
}
