package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/midepeter/train-ticket/database/models"
	"github.com/midepeter/train-ticket/database/postgres"
	"github.com/midepeter/train-ticket/server"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	db, err := postgres.New(&postgres.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBUrl:    os.Getenv("DB_URL"),
	})

	if err != nil {
		log.Fatal("failed to connect to postgresql database", err)
	}

	err = postgres.SetupDatabase(db,
		&models.Ticket{},
		&models.Station{},
		&models.Location{},
		&models.User{},
	)

	if err != nil {
		log.Fatal("failed to setup tables", err)
	}

	r := server.StartServer()

	r.RUN(":8000")
}
