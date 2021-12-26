package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mali3days/memberclub/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/members", handlers.GetAllMembers).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
