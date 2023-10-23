package main

import (
	"fmt"
	"net/http"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getBookings")
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createBooking")
}
