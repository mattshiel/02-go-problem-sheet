// Author: Matthew Shiel
// Date: 22/10/2017
// 03ServePage
// Info on Go documentation from https://golang.org/doc/articles/wiki/#tmp_3
// Adapted from: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html

package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type message struct {
	Message string
}

//
func handler(w http.ResponseWriter, r *http.Request) {
	// Serve the html file instead of hardcoding
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Call Seed using nanoseconds for a different random int every execution
	rand.Seed(int64(time.Now().Nanosecond()))

	// Check if the cookie 'target' is set and error handle if not
	if _, err := r.Cookie("target"); err != nil {
		//if cookie is not set generate a random number between 1 and 20
		// Limit random number between 1 and 20
		var randNum = ((rand.Int() % 19) + 1)

		// Convert random number to a string
		s := strconv.Itoa(randNum)

		cookie := http.Cookie{
			Name:  "target",
			Value: s,
		} // Create cookie

		// Set cookie target as new value of target
		http.SetCookie(w, &cookie)
	}
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
