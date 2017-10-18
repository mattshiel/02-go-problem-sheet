// Author: Matthew Shiel
// Date: 18/10/2017
// 01GuessingGame
// Adapted from: https://codegangsta.gitbooks.io/building-web-apps-with-go/content/creating_a_basic_web_app/index.html
// Info on Go documentation from https://golang.org/doc/articles/wiki/#tmp_3

package main

import (
	"fmt"
	"net/http"
)

// 
func handler(w http.ResponseWriter, r *http.Request) {
	// Response body - print Guessing Game
	fmt.Fprintln(w, "Guessing Game") 
}

func main() {
	// handler handles HTTP requests with the web root ("/") 
	http.HandleFunc("/", handler)

	// ListenAndServe will start the server and instruct it to listen on port 8080
    http.ListenAndServe(":8080", nil)
}