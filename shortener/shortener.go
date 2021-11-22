package shortener

import (
	"database/sql"
	"math/rand"
	"strings"

	_ "github.com/lib/pq"
)

type storage struct {
	m   map[string]string
	mb  map[string]string
	dbc *sql.DB
	db  string
}

var s = storage{nil, nil, nil, ""}

// Is called at the start of execution to choose the prefered data storage method and initialize it
func Init(c string) {
	switch c {
	case "m":
		s.m = make(map[string]string)
		s.mb = make(map[string]string)
		s.db = "memory"
	case "db":
		connStr := "user=postgres dbname=postgres password=test host=localhost sslmode=disable"
		var err error
		s.dbc, err = sql.Open("postgres", connStr)
		if err != nil {
			s.dbc = nil
		}
		s.db = "postgres"
	}
}

// Calls the correct save function for the selected database type and returns the result (the shortened url)
func SaveUrl(url string) string {
	var shorturl string
	switch s.db {
	case "memory":
		shorturl = SaveUrlMemory(url)
	case "postgres":
		shorturl = SaveUrlPostgres(url)
	}
	return shorturl
}

// Takes a full url, calls generateUrl, saves shortened url associated with full url and returns shortened url
func SaveUrlMemory(url string) string {
	shorturl := generateUrl()
	if s.m[shorturl] == url {
		return shorturl
	} else if s.mb[url] != "" {
		shorturl = s.mb[url]
		return shorturl
	} else if s.m[shorturl] != url && s.m[shorturl] != "" {
		SaveUrl(url)
	} else {
		s.m[shorturl] = url
		s.mb[url] = shorturl
	}
	return shorturl
}

func SaveUrlPostgres(url string) string {
	var shorturl string
	return shorturl
}

// Symbols that can be used to generate a shortened url
const availableSymbols string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// Generates short url of length 10 from the available symbols
func generateUrl() string {
	var b strings.Builder
	for i := 0; i < 10; i++ {
		rnd := []byte{availableSymbols[rand.Intn(63)]}
		b.Write(rnd)
	}
	return b.String()
}

// Takes shortened url as an argument, checks if it exists in the db and returns the asociated full url
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
