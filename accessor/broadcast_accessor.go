package accessor

import (
	"log"

	"github.com/Gokul-G/Remote-Download-Server/models"
	"github.com/gorilla/websocket"
)

type BroadcastMessage struct {
	DownloadIem    models.DownloadItem `json:"download_item"`
	DownloadedSize int64               `json:"downloaded_size"`
	Progress       float64             `json:"progress"`
}

var Clients = make(map[*websocket.Conn]bool)
var BroadcastChannel = make(chan BroadcastMessage)

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
