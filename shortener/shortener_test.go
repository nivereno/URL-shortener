package shortener

import "testing"

func TestSaveUrlMemory(t *testing.T) {
	input := []string{"https://golang.org/doc/", "https://golang.org/doc/", "https://golang.org/doc/12321321", "https://golang.org/doc/asdasdasd", "https://golang.org/doc/a213123213sdasdasd", "https://golang.org/doc/asda12312312321sdasd"}
	Init("memory")
	for _, v := range input {
		out := SaveUrl(v)
		if len(out) != 10 {
			t.Errorf("Shortened url is the wrong length")
		}
		if storage.m[out] != (v) {
			t.Errorf("Url not saved properly")
		}
	}
	a1 := SaveUrl("a")
	c1 := SaveUrl("c")
	a2 := SaveUrl("a")
	b1 := SaveUrl("b")
	c2 := SaveUrl("c")
	b2 := SaveUrl("b")
	c3 := SaveUrl("c")
	b3 := SaveUrl("b")
	b4 := SaveUrl("b")
	if a1 != a2 || b1 != b2 || b1 != b3 || b1 != b4 || c1 != c2 || c2 != c3 {
		t.Errorf("Saving the same data doesn't return the same key")
	}
}

func TestSaveUrlPostgres(t *testing.T) {
	input := []string{"https://golang.org/doc/", "https://golang.org/doc/", "https://golang.org/doc/12321321", "https://golang.org/doc/asdasd123asd", "https://golang.org/doc/a213123213sdasdasd", "https://golang.org/doc/asda12312312321sdasd", "https://golang.org/doc/qwertyu", "https://golang.org/doc/lglfdgdfg", "https://golang.org/doc/poiunnk", "https://golang.org/doc/poiunnkasdasd"}
	Init("postgres")
	for _, v := range input {
		out := SaveUrl(v)
		if len(out) != 10 {
			t.Errorf("Shortened url is the wrong length")
		}
	}
}

func TestGetUrlMemory(t *testing.T) {
	Init("memory")
	url := LookupUrl(SaveUrl("https://golang.org/doc/tutorial"))

	if url != "https://golang.org/doc/tutorial" {
		t.Errorf("Returned wrong url or no url")
	}
}

func TestGetUrlPostgres(t *testing.T) {
	Init("postgres")
	url := LookupUrl(SaveUrl("https://golang.org/doc/tutorial"))

	if url != "https://golang.org/doc/tutorial" {
		t.Errorf("Returned wrong url or no url")
	}
}
