package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRaceWebsites(t *testing.T) {
	slowServer := makeDelayedServer(5 * time.Millisecond)
	defer slowServer.Close()
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := RaceWebsites(slowURL, fastURL)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
