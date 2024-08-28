package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://asdf.asdf"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"waat://asdf.asdf",
	}

	want := map[string]bool{
		"http://google.com": true,
		"waat://asdf.asdf":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
