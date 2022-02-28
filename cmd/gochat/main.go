package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/bosamatheus/gochat/internal/api/handler"
	"github.com/bosamatheus/gochat/internal/infrastructure/repository"
	"github.com/bosamatheus/gochat/internal/usecase/chat"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	err := godotenv.Load("configs/.env." + env)
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// repository
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	repo := repository.NewChatRedis(ctx, client)

	// service
	service := chat.NewService(repo)

	// handler
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	handler := handler.NewHandler(upgrader, service)

	go handler.Broadcast()

	// router
	router := gin.Default()
	router.Static("/chat", "web/static")
	router.GET("/ws", func(c *gin.Context) {
		handler.ServeWebSocket(c.Writer, c.Request)
	})

	port := os.Getenv("SERVER_PORT")
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("error running server")
	}
}
