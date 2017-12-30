package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gokul-G/Remote-Download-Server/accessor"
	"github.com/Gokul-G/Remote-Download-Server/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetDownloadList(w http.ResponseWriter, r *http.Request) {

	downloads := accessor.GetDownloadListFromDB()
	response, _ := json.Marshal(downloads)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func StartDownload(w http.ResponseWriter, r *http.Request) {

	var downloadData models.DownloadData
	json.NewDecoder(r.Body).Decode(&downloadData)
	go download(&downloadData)
	// go download(&downloadData, broadcast)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SocketConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	accessor.Clients[ws] = true
	//Temp Fix
	for {

	}
	defer ws.Close()
}
