package main

import (
	"github.com/nivereno/URL-shortener/shortener"
)

func main() {
	c := "mem" //get this from arguments
	shortener.Init(c)
}
