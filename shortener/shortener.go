package shortener

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type storageStruct struct {
	// Map for saving data in memory
	memory map[string]string
	// Map but backwards! Used to quickly look up the shorturl in case the url has already been saved
	mBackwards map[string]string
	// The database connection
	databaseConnection *sql.DB
	// Used to specify which data storage solution is in use
	dbType string
}

var storage = storageStruct{nil, nil, nil, ""}

// Is called at the start of execution to choose the prefered data storage method and initialize it
func Init(mode string) {
	switch mode {
	case "memory":
		storage.memory = make(map[string]string)
		storage.mBackwards = make(map[string]string)
		storage.dbType = "memory"
	case "postgres":
		connStr := "user=postgres dbname=postgres password=test host=docker-postgres sslmode=disable"
		storage.databaseConnection, _ = sql.Open("postgres", connStr)
		err := storage.databaseConnection.Ping()
		if err != nil {
			panic(err)
		}
		storage.dbType = "postgres"
	default:
		panic("No or unsupported database selected, please provide the env -e storage=(either postgres or memory) variable when using docker compose")
	}
}

// Calls the correct save function for the selected database type and returns the result (the shortened url)
func SaveUrl(url string) string {
	var shorturl string
	switch storage.dbType {
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
	if storage.mBackwards[url] != "" {
		shorturl = storage.mBackwards[url]
		return shorturl
	} else if storage.memory[shorturl] != url && storage.memory[shorturl] != "" {
		saveUrlMemory(url)
	} else {
		storage.memory[shorturl] = url
		storage.mBackwards[url] = shorturl
	}
	return shorturl
}

// Takes a full url, calls generateUrl, saves shortened url associated with full url in a postgres sql database and returns shortened url
func saveUrlPostgres(url string) string {
	shorturl := checkFullUrl(url)
	if shorturl == "" {
		for {
			shorturl = generateUrl()
			if !checkCollision(shorturl) {
				query := `INSERT INTO urls (shortenedurl, fullurl) VALUES ($1, $2)`
				_, err := storage.databaseConnection.Exec(query, shorturl, url)
				if err != nil {
					shorturl = "Failed to insert into DB"
				}
				break
			}
		}
	}
	return shorturl
}

// Checks if the full url already exists in the database and returns the correct shortened url if it does
func checkFullUrl(url string) string {
	query := `SELECT shortenedurl FROM urls WHERE fullurl=$1`
	row := storage.databaseConnection.QueryRow(query, url)
	var shorturl string
	row.Scan(&shorturl)
	return shorturl
}

// Checks if the generated shortened url has been used before
func checkCollision(shorturl string) bool {
	query := `SELECT fullurl FROM urls WHERE shortenedurl=$1`
	row := storage.databaseConnection.QueryRow(query, shorturl)
	shortCollision := false
	var fullurl string
	row.Scan(&fullurl)
	if fullurl != "" {
		shortCollision = true
	}
	return shortCollision
}

// Symbols that can be used to generate a shortened url
const availableSymbols string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// Generates short url of length 10 from the available symbols
func generateUrl() string {
	rand.Seed(time.Now().UnixNano())
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
	switch storage.dbType {
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
	if storage.memory[shorturl] == "" {
		fullurl = "Url does not exist"
	} else {
		fullurl = storage.memory[shorturl]
	}
	return fullurl
}

// Takes shortened url as an argument, checks if it exists in postgres dbType and returns the asociated full url
func lookupUrlPostgres(shorturl string) string {
	query := `SELECT fullurl FROM urls WHERE shortenedurl=$1`
	row := storage.databaseConnection.QueryRow(query, shorturl)
	var fullurl string
	row.Scan(&fullurl)
	if fullurl == "" {
		fullurl = "Url does not exist"
	}
	return fullurl
}
