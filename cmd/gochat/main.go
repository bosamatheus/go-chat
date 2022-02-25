package main

import (
	"log"
	"os"

	"github.com/bosamatheus/gochat/internal/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	err := godotenv.Load("configs/.env." + env)
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	r := gin.Default()

	r.GET("/api/v1/ping", handler.Ping)

	port := os.Getenv("PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatal("error running server")
	}
}
