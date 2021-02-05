package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/midepeter/train-ticket/database/postgres"
	"github.com/midepeter/train-ticket/handler"
)

func main(){
	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	db, err := postgres.New(&postgres.Config{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DBUrl: os.Getenv("DB_URL"),
	})

	if err != nil {
		log.Fatal("failed to connect to postgresql database", err)
	}

	err := postgres.SetupDatabase(db,
		&models.Ticket,
		&models.Station,
		&models.Location,
		&models.User{}
	)

	if err != nil {
		log.Fatal("failed to setup tables", err)
	}

	r := gin.Default()

	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)
	r.GET("/bookings", GetAllTicket)
	r.POST("/tickets", CreateTicket)
	r.PUT("/tickets/:id", UpdateTicket)
	r.GET("/tickets/:id", GetTicket)
	r.DELETE("/tickets/:id", DeleteTicket)

	r.RUN(":8000")
}