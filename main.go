package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nivereno/URL-shortener/shortener"
)

func main() {
	c := "mem" //get this from arguments
	shortener.Init(c)
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("sdsafdsafdfasfko")
	shortener.SaveUrl("1342dasfasdf")
	shortener.SaveUrl("aksndoakwndo")
	println(shortener.LookupUrl("a21321ksndoa12312kwsndo"))
	println(shortener.LookupUrl("sdadkmwkmko"))

	l := log.New(os.Stdout, "shortener-api", log.LstdFlags)
	r := chi.NewRouter()

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
