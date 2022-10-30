package iteration

import "testing"

func TestRepeat(t *testing.T) {
	sum := Repeat("a")
	expected := "aaaaa"
	if sum != expected {
		t.Errorf("expected %q, got %q", expected, sum)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
