package accessor

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var Clients = make(map[*websocket.Conn]bool) // connected clients
var BroadcastChannel = make(chan Message)    // broadcast channel

func HandleBroadcast() {
	go func() {
		for {
			msg := <-BroadcastChannel
			for client := range Clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(Clients, client)
				}
			}
		}
	}()
}
