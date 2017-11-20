package controllers

import (
	"encoding/json"
	"net/http"

	"../accessor"
)

func GetDownloadList(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if r.Method == "OPTIONS" {
		return
	}

	downloads := accessor.GetDownloadListFromDB()
	response, _ := json.Marshal(downloads)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
