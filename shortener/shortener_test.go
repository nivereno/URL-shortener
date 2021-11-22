package shortener

import "testing"

func TestSaveUrlMemory(t *testing.T) {
	input := []string{"https://golang.org/doc/", "https://golang.org/doc/", "https://golang.org/doc/12321321", "https://golang.org/doc/asdasdasd"}
	Init("m")
	for _, v := range input {
		out := SaveUrl(v)
		if len(out) != 10 {
			t.Errorf("Shortened url is the wrong length")
		}
		if s.m[out] != (v) {
			t.Errorf("Url not saved properly")
		}
	}
}

func TestSaveUrlPostgres(t *testing.T) {
	input := []string{"https://golang.org/doc/", "https://golang.org/doc/", "https://golang.org/doc/12321321", "https://golang.org/doc/asdasdasd"}
	Init("db")
	for _, v := range input {
		out := SaveUrl(v)
		if len(out) != 10 {
			t.Errorf("Shortened url is the wrong length")
		}
	}
}

func TestGetUrlMemory(t *testing.T) {
	Init("m")
	url := LookupUrl(SaveUrl("https://golang.org/doc/tutorial"))

	if url != "https://golang.org/doc/tutorial" {
		t.Errorf("Returned wrong url or no url")
	}
}

func TestGetUrlPostgres(t *testing.T) {
	Init("db")
	url := LookupUrl(SaveUrl("https://golang.org/doc/tutorial"))

	if url != "https://golang.org/doc/tutorial" {
		t.Errorf("Returned wrong url or no url")
	}
}
