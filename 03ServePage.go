// Author: Matthew Shiel
// Date: 18/10/2017
// 03ServePage
// Adapted from http://www.alexedwards.net/blog/serving-static-sites-with-go

package main

import (
	"log"
	"net/http"
)

func main() {
	 // Creates the handler to respond to the HTTP request
	fs := http.FileServer(http.Dir("/02-go-problem-sheet/index.html"))
	// Handles HTTP requests with the web root ("/") 
	http.Handle("/", fs)
  
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}