package main

import (
	"github.com/nivereno/URL-shortener/shortener"
)

func main() {
	c := "mem" //get this from arguments
	shortener.Init(c)
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("aksndoakwndo")
	println(shortener.SaveUrl("13213v23"))

}
