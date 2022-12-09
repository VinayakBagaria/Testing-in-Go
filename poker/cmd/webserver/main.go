package poker

import (
	"log"
	"net/http"
	"os"
	poker "testing-go/poker"
)

const dbFileName = "game.db.json"

func StartPlayerServer() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
