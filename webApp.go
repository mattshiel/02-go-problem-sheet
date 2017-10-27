// Author: Matthew Shiel
// Date: 22/10/2017
// webApp
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
	Guess   string
	Win     bool
}

//
func handler(w http.ResponseWriter, r *http.Request) {
	// Serve the html file instead of hardcoding
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form in /guess
	r.ParseForm()

	// Get the user guess and store it in a variable to be returned, if none the text won't be displayed
	guess := r.FormValue("guess")

	// Ask user to guess a number
	question := "Guess a number between 1 and 20: "

	// Call Seed using nanoseconds for a different random int every execution
	rand.Seed(int64(time.Now().Nanosecond()))

	// Declare variable before error handling so it is within scope to be compared to guess
	cookie, err := r.Cookie("target"); 

	// Check if the cookie 'target' is set and error handle if not
	if err != nil {
		//if cookie is not set generate a random number between 1 and 20
		// Limit random number between 1 and 20
		var randNum = ((rand.Int() % 19) + 1)

		// Convert random number to a string
		s := strconv.Itoa(randNum)

		// Create cookie
		cookie := http.Cookie{
			Name:  "target",
			Value: s,
		}

		// Set cookie target as new value of target
		http.SetCookie(w, &cookie)
	}

	guessInt, _ := strconv.Atoi(guess)
	targetInt, _ := strconv.Atoi(cookie.Value)

	// Set the message and guess Strings and Win boolean values
	m := message{Message: question, Guess: guess, Win: false}

	// Compare target and guess 
	if targetInt == guessInt {
		// message = ("Congratulations you guessed correctly! You win!")
	} else if targetInt < guessInt {
		//message = ("Lower! Your guess was too high!")
	} else {
		//message = ("Higher! Your guess was too low!")
	}

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
	http.ListenAndServe(":8089", nil)
}
