package main

import (
	"fmt"
	"net/http"
	"os"

	"./accessor"
	"./routers"
	"github.com/codegangsta/negroni"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Println("Remote Download Server")
	fmt.Println("Server Started...")

	accessor.InitDB()

	router := routers.InitRoutes()
	router.HandleFunc("/", serverHandler)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + port)

}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "")
}