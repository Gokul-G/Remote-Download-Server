package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Gokul-G/Remote-Download-Server/accessor"
	"github.com/Gokul-G/Remote-Download-Server/models"
)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
