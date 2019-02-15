package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chapin666/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
