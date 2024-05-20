package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Error upgrading to WebSocket: %s", err))
		return
	}
	defer ws.Close()
	counter := 0
	for {
		counter++
		ws.WriteJSON(map[string]int{"counter": counter})
		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocketConnection)
	panic(http.ListenAndServe(":8080", nil))
}
