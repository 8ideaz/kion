package server

import (
	"log"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("./packages"))
	http.Handle("/packages/", http.StripPrefix("/packages", fs))

	log.Println("Serving repository on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
