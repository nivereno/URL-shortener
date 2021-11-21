package shortener

import (
	"math/rand"
	"strings"
)

type storage struct {
	m  map[string]string
	d  string
	db string
}

var s = storage{nil, "", ""}

func Init(c string) {
	switch c {
	case "m":
		s.m = make(map[string]string)
		s.db = "memory"
	case "db":
		s.d = "db connection blah blah blah"
		s.db = "postgres"
	}
}

func SaveUrl(url string) string {
	shorturl := generateUrl()
	switch s.db {
	case "memory":
		for k, v := range s.m {
			if v == url {
				shorturl = k
			} else if k == shorturl {
				SaveUrl(url)
			}
		}
		s.m[shorturl] = url

	case "postgres":
	}
	return shorturl
}

const availableSymbols string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func generateUrl() string {
	var b strings.Builder
	for i := 0; i < 10; i++ {
		rnd := []byte{availableSymbols[rand.Intn(63)]}
		b.Write(rnd)
	}
	return b.String()
}

func LookupUrl(shorturl string) string {
	var fullurl string
	switch s.db {
	case "memory":
		if s.m[shorturl] == "" {
			fullurl = "Url does not exist"
		} else {
			fullurl = s.m[shorturl]
		}
	case "postgres":
	}
	return fullurl
}
