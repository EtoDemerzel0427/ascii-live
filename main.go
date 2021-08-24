package main

import (
	"net/http"
	"github.com/EtoDemerzel0427/ascii-live/backend"
)


func main() {
	backend.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}

