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

func (s *Shortener) PostUrl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST url")

	url := r.URL.Query().Get("url")
	code, err := rw.Write([]byte(shortener.SaveUrl(url)))
	if err != nil {
		http.Error(rw, "POST failed", code)
	}
}

func (s *Shortener) GetUrl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle GET url")

	shorturl := r.URL.Path[1:]
	code, err := rw.Write([]byte(shortener.LookupUrl(shorturl)))
	if err != nil {
		http.Error(rw, "GET failed", code)
	}
}
