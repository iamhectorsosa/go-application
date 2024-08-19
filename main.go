package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8080", server))
}
