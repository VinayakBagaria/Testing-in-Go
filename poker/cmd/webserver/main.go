package poker

import (
	"log"
	"net/http"
	poker "testing-go/poker"
)

const dbFileName = "game.db.json"

func StartPlayerServer() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server, err := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
