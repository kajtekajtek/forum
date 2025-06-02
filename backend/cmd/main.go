package main

import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	"github.com/kajtekajtek/forum/backend/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open postgres database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("database not reachable: %v", err)
	}

	log.Println("connected to postgresql!")
}
