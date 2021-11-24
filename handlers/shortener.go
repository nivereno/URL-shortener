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

// Handles a post request, takes the full url from data (-d url="some data"), calls save url and writes the newly generated shorturl as a response
func (s *Shortener) PostUrl(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	s.l.Println("Handle POST url with data -d " + url)
	code, err := rw.Write([]byte(shortener.SaveUrl(url)))
	if err != nil {
		http.Error(rw, "POST failed", code)
	}
}

// Handles a get request, grabs the short url from the full call, calls lookup url to find the associated full url and return it
func (s *Shortener) GetUrl(rw http.ResponseWriter, r *http.Request) {
	shorturl := r.URL.Path[1:]
	s.l.Println("Handle GET url for shortened url " + shorturl)
	code, err := rw.Write([]byte(shortener.LookupUrl(shorturl)))
	if err != nil {
		http.Error(rw, "GET failed", code)
	}
}
