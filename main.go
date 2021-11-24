package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/nivereno/URL-shortener/handlers"
	"github.com/nivereno/URL-shortener/shortener"
)

func main() {
	l := log.New(os.Stdout, "shortener-api", log.LstdFlags)
	r := mux.NewRouter()
	sh := handlers.NewShortener(l)

	if os.Getenv("storage") == "postgres" {
		fmt.Println("Waiting for the postgres container to start up")
		time.Sleep(10 * time.Second)
		fmt.Println("Done")
	}
	c := os.Getenv("storage")
	shortener.Init(c)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", sh.PostUrl)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{shorturl:[a-zA-Z0-9_]+}", sh.GetUrl)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Printf("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
	os.Exit(0)
}
