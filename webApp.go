// Author: Matthew Shiel
// Date: 22/10/2017
// 03ServePage
// Info on Go documentation from https://golang.org/doc/articles/wiki/#tmp_3
// Adapted from: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html

package main

import (
	"html/template"
	"net/http"
)

type  message struct {
	Message string 
}

//
func handler(w http.ResponseWriter, r *http.Request) {
	// Serve the html file instead of hardcoding
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Set Message
	m := message{Message: "Guess a number between 1 and 20: "}
	// Parse template guess.tmpl
	t, _ := template.ParseFiles("guess.tmpl") 
	// Applies parsed template 't' and writes to output writer
	t.Execute(w, m)
}

func main() {
	// handler handles HTTP requests with the web root ("/")
	http.HandleFunc("/", handler)

	// guessHandler handles HTTP requests with the guess root("/guess")
	http.HandleFunc("/guess", guessHandler)

	// ListenAndServe will start the server and instruct it to listen on port 8080
	http.ListenAndServe(":8080", nil)
}
