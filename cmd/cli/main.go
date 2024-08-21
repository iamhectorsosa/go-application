package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/iamhectorsosa/go-application"
)

const dbFileName = "game.db.json"

func main() {
	store, cleanStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	defer cleanStore()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
