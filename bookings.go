package main

import (
	"fmt"
	"net/http"
)

func getBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getBookings")
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createBooking")
}
