package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bosamatheus/gochat/internal/api/presenter"
	"github.com/bosamatheus/gochat/internal/usecase/chat"
	"github.com/gorilla/websocket"
)

const chatKey = "chat"

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan presenter.ChatMessage)

type Handler struct {
	upgrader websocket.Upgrader
	service  chat.UseCase
}

func NewHandler(upgrader websocket.Upgrader, service chat.UseCase) *Handler {
	return &Handler{
		upgrader: upgrader,
		service:  service,
	}
}

func (h *Handler) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to set websocket upgrade: %s", err.Error())
		return
	}
	defer conn.Close()
	clients[conn] = true

	if h.service.ChatExists(chatKey) {
		h.sendChatHistory(conn)
	}

	for {
		var msg presenter.ChatMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			break
		}
		// send new message to the channel
		broadcaster <- msg
	}
}

func (h *Handler) Broadcast() {
	for {
		// grab any next message from channel
		msg := <-broadcaster
		json, err := json.Marshal(msg)
		if err != nil {
			log.Printf("failed to marshal chat message: %s", err.Error())
			return
		}

		err = h.service.SaveMessage(chatKey, json)
		if err != nil {
			log.Printf("failed to save chat message: %s", err.Error())
			return
		}
		h.notifyAll(msg)
	}
}

func (h *Handler) sendChatHistory(conn *websocket.Conn) {
	messages, err := h.service.GetChatHistory(chatKey)
	if err != nil {
		log.Printf("failed to get previous chat messages: %s", err.Error())
	}

	// send previous messages
	for _, message := range messages {
		var msg presenter.ChatMessage
		err = json.Unmarshal([]byte(message), &msg)
		if err != nil {
			log.Printf("failed to unmarshal chat message: %s", err.Error())
			return
		}
		h.sendToClient(conn, msg)
	}
}

func (h *Handler) notifyAll(msg presenter.ChatMessage) {
	// send to every client currently connected
	for client := range clients {
		h.sendToClient(client, msg)
	}
}

func (h *Handler) sendToClient(conn *websocket.Conn, msg presenter.ChatMessage) {
	err := conn.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		conn.Close()
		delete(clients, conn)
	}
}

func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}
