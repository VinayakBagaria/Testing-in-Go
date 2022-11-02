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
	t.Run("test fast server", func(t *testing.T) {
		slowServer := makeDelayedServer(10 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := RaceWebsites(slowURL, fastURL, 4*time.Millisecond)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(15 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayedServer(10 * time.Millisecond)
		defer fastServer.Close()

		_, err := RaceWebsites(slowServer.URL, fastServer.URL, 4*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
