package handler

import (
	"net/http"

	"github.com/bosamatheus/gochat/internal/usecase/chat"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service chat.UseCase
}

func NewHandler(service chat.UseCase) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	pong, err := h.service.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": pong})
}
