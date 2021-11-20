package shortener

import "testing"

func TestSaveUrl(t *testing.T) {
	Init("mem")
	out := SaveUrl("https://golang.org/doc/")

	if len(out) != 10 {
		t.Errorf("Shortened url is the wrong length")
	}
	if s.m[out] != ("https://golang.org/doc/") {
		t.Errorf("Url not saved properly")
	}

	//add test for db
	Init("db")
}

func TestGetUrl(t *testing.T) {
	Init("mem")
	url := LookupUrl(SaveUrl("https://golang.org/doc/tutorial"))

	if url != "https://golang.org/doc/tutorial" {
		t.Errorf("Returned wrong url or no url")
	}

	Init("db")
}
