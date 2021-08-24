package main

import "github.com/EtoDemerzel0427/ascii-live/backend"

func main() {
	//backend.SetupRoutes()
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	return
	//}
	t := backend.NewTerminal()
	t.Debug()
}
