package handlers

import "log"

type Shortener struct {
	l *log.Logger
}

func NewShortener(l *log.Logger) *Shortener {
	return &Shortener{l}
}
