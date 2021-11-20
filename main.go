package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/nivereno/URL-shortener/handlers"
	"github.com/nivereno/URL-shortener/shortener"
)

func main() {
	c := "mem" //get this from arguments
	shortener.Init(c)

	l := log.New(os.Stdout, "shortener-api", log.LstdFlags)
	r := mux.NewRouter()
	sh := handlers.NewShortener(l)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", sh.AddUrl)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", sh.LookupUrl)

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
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	cancel()
}
