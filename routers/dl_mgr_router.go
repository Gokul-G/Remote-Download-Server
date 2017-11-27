package routers

import (
	"github.com/Gokul-G/Remote-Download-Server/controllers"
	"github.com/gorilla/mux"
)

func SetDownloadListRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/downloads", controllers.GetDownloadList).Methods("GET")
	return router
}

func SetStartDownloadRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/downloads", controllers.StartDownload).Methods("POST")
	return router
}
