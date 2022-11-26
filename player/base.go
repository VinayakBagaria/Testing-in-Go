package player

import (
	"log"
	"net/http"
)

func StartPlayerServer() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
