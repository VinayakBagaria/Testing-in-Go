package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://twitter.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"https://google.com":      true,
		"https://twitter.com":     true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebsite, websites)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(10 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
