package main

import (
	c "aqbiluts/controller"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // Connection
	"github.com/gorilla/mux"           // Router
	"github.com/rs/cors"
)

func main() {
	// Start Mux
	router := mux.NewRouter()

	// General Endpoint
	router.HandleFunc("/rooms", c.GetAllRooms).Methods("GET")
	router.HandleFunc("/room", c.GetDetailRoom).Methods("GET")

	// CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	// Start Server
	http.Handle("/", router)

	log.Println("Starting " + os.Getenv("APP_NAME"))
	log.Println("Connected to port 8081")
	log.Fatal(http.ListenAndServe(":8081", corsHandler.Handler(router)))
}
