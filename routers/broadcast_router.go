package routers

import (
	"github.com/Gokul-G/Remote-Download-Server/controllers"
	"github.com/gorilla/mux"
)

func SetBroadcastRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ws", controllers.SocketConnection).Methods("GET")
	return router
}
