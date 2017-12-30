package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetDownloadListRoutes(router)
	router = SetStartDownloadRoutes(router)
	router = SetBroadcastRoutes(router)

	return router
}
