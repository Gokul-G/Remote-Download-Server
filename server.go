package main

import (
	"fmt"
	"os"

	"github.com/Gokul-G/Remote-Download-Server/accessor"
	"github.com/Gokul-G/Remote-Download-Server/routers"
	"github.com/codegangsta/negroni"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Remote Download Server")
	fmt.Println("Server Started...")

	//Init DB
	accessor.InitDB()

	//Init Broadcast
	accessor.HandleBroadcast()

	//Init Routes
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + port)
}
