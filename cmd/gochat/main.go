package main

import (
	"context"
	"log"
	"os"

	"github.com/bosamatheus/gochat/internal/api/handler"
	"github.com/bosamatheus/gochat/internal/infrastructure/repository"
	"github.com/bosamatheus/gochat/internal/usecase/chat"
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
	ctx := context.Background()

	repo := repository.NewChatRedis(&ctx)
	service := chat.NewService(repo)
	handler := handler.NewHandler(service)

	r := gin.Default()
	v1 := r.Group(os.Getenv("API_V1"))
	v1.GET("/ping", handler.Ping)

	port := os.Getenv("SERVER_PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatal("error running server")
	}
}
