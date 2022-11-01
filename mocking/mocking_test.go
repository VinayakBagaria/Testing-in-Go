package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	calls []string
}

func (s *SpyCountdownOperations) Write(p []byte) (int, error) {
	s.calls = append(s.calls, write)
	return 0, nil
}

func (s *SpyCountdownOperations) Sleep() {
	s.calls = append(s.calls, sleep)
}

func TestCountdown(t *testing.T) {
	t.Run("test before every print", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Countdown(&buffer, &SpyCountdownOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})

	t.Run("test before every print", func(t *testing.T) {
		spySleeper := &SpyCountdownOperations{}
		Countdown(spySleeper, spySleeper)
		want := []string{write, sleep, write, sleep, write, sleep, write}
		if !reflect.DeepEqual(want, spySleeper.calls) {
			t.Errorf("wanted %v, got %v", want, spySleeper.calls)
		}
	})
}
