package syncer

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter X times leaves it at X", func(t *testing.T) {
		want := 3
		counter := Counter{}
		for i := 0; i < want; i++ {
			counter.Inc()
		}
		assertCounter(t, &counter, want)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		want := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, &counter, want)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
