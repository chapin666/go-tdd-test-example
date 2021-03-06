package main

import (
	"log"
	"net/http"

	"github.com/chapin666/poker"
)

const dbFileName = "game.db.json"

func main() {

	store, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := poker.NewPlayerServer(store)
	//server := NewPlayerServer(NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
