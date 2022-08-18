package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marutijogdand17/golang-bankapp/routes"
)

func main() {

	router := routes.NewRouter()
	routes.ConfigureRoutes(router)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	fmt.Println("Server started, Listening on port: 8000")
	if err := server.ListenAndServe(); err != nil {
		log.Println("Error while starting the server", err.Error())
	}
}
