package shortener

import (
	"database/sql"
	"math/rand"
	"strings"

	_ "github.com/lib/pq"
)

type storageStruct struct {
	// Map for saving data in memory
	m map[string]string
	// Map but backwards! Used to quickly look up the shorturl in case the url has already been saved
	mb map[string]string
	// The database connection
	dbc *sql.DB
	// Used to specify which data storage solution is in use
	db string
}

var storage = storageStruct{nil, nil, nil, ""}

// Is called at the start of execution to choose the prefered data storage method and initialize it
func Init(c string) {
	switch c {
	case "m":
		storage.m = make(map[string]string)
		storage.mb = make(map[string]string)
		storage.db = "memory"
	case "db":
		connStr := "user=postgres dbname=postgres password=test host=localhost sslmode=disable"
		var err error
		storage.dbc, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		err = storage.dbc.Ping()
		if err != nil {
			panic(err)
		}
		storage.db = "postgres"
	}
}

// Calls the correct save function for the selected database type and returns the result (the shortened url)
func SaveUrl(url string) string {
	var shorturl string
	switch storage.db {
	case "memory":
		shorturl = saveUrlMemory(url)
	case "postgres":
		shorturl = saveUrlPostgres(url)
	default:
		shorturl = "No database selected"
	}
	return shorturl
}

// Takes a full url, calls generateUrl, saves shortened url associated with full url in memory and returns shortened url
func saveUrlMemory(url string) string {
	shorturl := generateUrl()
	if storage.mb[url] != "" {
		shorturl = storage.mb[url]
		return shorturl
	} else if storage.m[shorturl] != url && storage.m[shorturl] != "" {
		saveUrlMemory(url)
	} else {
		storage.m[shorturl] = url
		storage.mb[url] = shorturl
	}
	return shorturl
}

// Takes a full url, calls generateUrl, saves shortened url associated with full url in a postgres sql database and returns shortened url
func saveUrlPostgres(url string) string {
	shorturl := generateUrl()
	check, shortCollision := saveUrlPostgresHelper(url, shorturl)
	if check != "" {
		shorturl = check
	} else if shortCollision {
		saveUrlPostgres(url)
	} else {
		query := `INSERT INTO urls (shortenedurl, fullurl) VALUES ($1, $2)`
		_, err := storage.dbc.Exec(query, shorturl, url)
		if err != nil {
			shorturl = "Failed to insert into DB"
		}
		return shorturl
	}
	return shorturl

}

func saveUrlPostgresHelper(url string, generatedurl string) (string, bool) {
	query := `SELECT shortenedurl FROM urls WHERE fullurl=$1`
	row := storage.dbc.QueryRow(query, url)
	var shorturl string
	row.Scan(&shorturl)

	query2 := `SELECT fullurl FROM urls WHERE shortenedurl=$1`
	row2 := storage.dbc.QueryRow(query2, generatedurl)
	shortCollision := false
	var fullurl string
	row2.Scan(&fullurl)
	if fullurl != url && fullurl != "" {
		shortCollision = true
	}
	return shorturl, shortCollision
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

// Calls the correct lookup function for the selected database type and returns the result (the full url)
func LookupUrl(shorturl string) string {
	var fullurl string
	switch storage.db {
	case "memory":
		fullurl = lookupUrlMemory(shorturl)
	case "postgres":
		fullurl = lookupUrlPostgres(shorturl)
	default:
		fullurl = "No database selected"
	}
	return fullurl
}

// Takes shortened url as an argument, checks if it exists in the in memory map and returns the asociated full url
func lookupUrlMemory(shorturl string) string {
	var fullurl string
	if storage.m[shorturl] == "" {
		fullurl = "Url does not exist"
	} else {
		fullurl = storage.m[shorturl]
	}
	return fullurl
}

// Takes shortened url as an argument, checks if it exists in postgres db and returns the asociated full url
func lookupUrlPostgres(shorturl string) string {
	query := `SELECT fullurl FROM urls WHERE shortenedurl=$1`
	row := storage.dbc.QueryRow(query, shorturl)
	var fullurl string
	row.Scan(&fullurl)
	return fullurl
}
