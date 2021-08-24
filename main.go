package main

import (
	"github.com/EtoDemerzel0427/ascii-live/backend"
	"net/http"
)

func main() {
	backend.SetupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return 
	}
}
