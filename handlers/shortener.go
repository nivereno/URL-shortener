package handlers

import (
	"log"
	"net/http"

	"github.com/nivereno/URL-shortener/shortener"
)

type Shortener struct {
	l *log.Logger
}

func NewShortener(l *log.Logger) *Shortener {
	return &Shortener{l}
}

// Handles a post request, takes the full url from data {url:}, calls save url and writes the newly generated shorturl as a response
func (s *Shortener) PostUrl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST url")

	r.ParseForm()
	url := r.Form.Get("url")
	code, err := rw.Write([]byte(shortener.SaveUrl(url)))
	if err != nil {
		http.Error(rw, "POST failed", code)
	}
}

// Handles a get request, grabs the short url from the full call, calls lookup url to find the associated full url and returns it
func (s *Shortener) GetUrl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle GET url")

	shorturl := r.URL.Path[1:]
	code, err := rw.Write([]byte(shortener.LookupUrl(shorturl)))
	if err != nil {
		http.Error(rw, "GET failed", code)
	}
}
