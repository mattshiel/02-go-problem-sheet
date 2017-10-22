// Author: Matthew Shiel
// Date: 18/10/2017
// 02MakeH1
// Info on Go documentation from https://golang.org/doc/articles/wiki/#tmp_3

package main

import (
	"fmt"
	"net/http"
)

//
func handler(w http.ResponseWriter, r *http.Request) {
	// Encapsulate Guessing Game in a <h1> tag
	title := "<h1>Guessing Game</h1>"

	// Print <h1> Guessing Game
	fmt.Fprintln(w, title)
}

func main() {
	// handler handles HTTP requests with the web root ("/")
	http.HandleFunc("/", handler)

	// ListenAndServe will start the server and instruct it to listen on port 8080
	http.ListenAndServe(":8080", nil)
}
