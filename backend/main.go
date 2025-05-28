package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func telemetryWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		data := map[string]interface{}{
			"carId": 44,
			"speed": 270 + time.Now().Unix()%10,
			"x":     123.45,
			"y":     678.90,
		}
		if err := conn.WriteJSON(data); err != nil {
			log.Println("Write error:", err)
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", telemetryWS)
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
