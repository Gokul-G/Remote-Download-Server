package routers

import (
	"../controllers"
	"github.com/gorilla/mux"
)

func SetDownloadListRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/downloads", controllers.GetDownloadList).Methods("GET")
	return router
}
