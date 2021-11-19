package main

import (
	"github.com/nivereno/URL-shortener/shortener"
)

func main() {
	c := "mem" //get this from arguments
	shortener.Init(c)
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("sdadkmwkmko")
	shortener.SaveUrl("aksndoakwndo")
	println(shortener.LookupUrl("a21321ksndoa12312kwsndo"))
	println(shortener.LookupUrl("sdadkmwkmko"))

}
