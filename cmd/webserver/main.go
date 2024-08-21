package main

import (
	"log"
	"net/http"

	poker "github.com/iamhectorsosa/go-application"
)

const dbFileName = "game.db.json"

func main() {
	store, cleanStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	defer cleanStore()

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8080", server))
}
