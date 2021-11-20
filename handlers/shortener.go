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

	url := "123"
	shortener.SaveUrl(url)
}

func (s *Shortener) GetUrl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle GET url")

	shorturl := "132"
	shortener.LookupUrl(shorturl)
}
