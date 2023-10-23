package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/api/bookings", getBookings).Methods("GET")
	router.HandleFunc("/api/bookings", createBooking).Methods("POST")

	// Iniciar el servidor en el puerto 9000
	fmt.Println("API server listening on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
