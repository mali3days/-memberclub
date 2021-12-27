package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mali3days/memberclub/pkg/handlers"
	"github.com/rs/cors"
)

func main() {
	l := log.New(os.Stdout, "memberclub-api ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewMembers(l)

	// create a new serve mux and register the handlers
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetMembers)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddMember)
	postRouter.Use(ph.MiddlewareValidateMember)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*",
		},
	})

	port := ":4000"
	fmt.Println("Server is running on port " + port)

	// create a new server
	s := http.Server{
		Addr:         port,                     // configure the bind address
		Handler:      corsOpts.Handler(router), // set the default handler
		ErrorLog:     l,                        // set the logger for the server
		ReadTimeout:  5 * time.Second,          // max time to read request from the client
		WriteTimeout: 10 * time.Second,         // max time to write response to the client
		IdleTimeout:  120 * time.Second,        // max time for connections using TCP Keep-Alive
	}

	err := s.ListenAndServe()

	if err != nil {
		l.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
