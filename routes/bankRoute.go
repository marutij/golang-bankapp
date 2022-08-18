package routes

import (
	"github.com/gorilla/mux"
	"github.com/marutijogdand17/golang-bankapp/handlers"
	"github.com/marutijogdand17/golang-bankapp/middleware"
)

//var router mux.Router

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func ConfigureRoutes(router *mux.Router) {

	router.HandleFunc("/", handlers.Welcome).Methods("GET")
	router.HandleFunc("/users/login", handlers.Login).Methods("POST")
	router.HandleFunc("/accounts", handlers.GetAccount).Methods("GET")
	router.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", handlers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", handlers.DeleteAccount).Methods("PUT")

	router.Use(middleware.LoggingMiddleware)
}
