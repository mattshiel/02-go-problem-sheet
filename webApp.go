// Author: Matthew Shiel
// Date: 22/10/2017
// 03ServePage
// Info on Go documentation from https://golang.org/doc/articles/wiki/#tmp_3
// 

package main

import (
	"net/http"
)

// 
func handler(w http.ResponseWriter, r *http.Request) {
	// Serve the html file instead of hardcoding
	http.ServeFile(w, r, "index.html") 
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the html file instead of hardcoding
	http.ServeFile(w, r, "guess.html") 
}


func main() {
	// handler handles HTTP requests with the web root ("/") 
	http.HandleFunc("/", handler)

	// guessHandler handles HTTP requests with the guess root("/") 
	http.HandleFunc("/", guessHandler)

	// ListenAndServe will start the server and instruct it to listen on port 8080
    http.ListenAndServe(":8080", nil)
}