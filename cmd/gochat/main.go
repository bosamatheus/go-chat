package main

import (
	"log"
	"os"

	"github.com/bosamatheus/gochat/internal/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	r := gin.Default()

	r.GET("/api/v1/ping", handler.Ping)

	port := os.Getenv("PORT")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("error running server")
	}
}
