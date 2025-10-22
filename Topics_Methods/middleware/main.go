package main

import (
	"fmt"
	"log"
	"net/http"
)

var database = map[string]string{
	"admin": "password",
}

// in Go middleware is like a wrapper in other languages - in go this is specifically for http requests
// generally used for eg. adding authentication, logging, request modification to endpoints in a modular way
// the below is a web API with only a /welcome endpoint that we want to protect by authentication

// This solution is not scalable as if there are more routes to take, it will get redundant as we would need to repeate the if statement in each endpoints function
// func authenticate(r *http.Request) bool {
// 	user := r.FormValue("user")
// 	password := r.FormValue("password")

// 	if pass, ok  := database[user]; !ok || pass != password {
// 		return false
// 	}

// 	return true
// }

// This is where middleware comes in - middleware is like a function wrapper
// middleware(fn function) -> function (receives a function, returns another)
// In MW case it is an http.HandlerFunc -> this type is just a function that looks like this: type HandlerFunc func(ResponseWriter, *Request)

func authMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		password := r.FormValue("password")

		// check if user exists and password matchs
		if pass, ok := database[user]; !ok || pass != password {
			err := http.StatusUnauthorized
			http.Error(w, "Invalid username or password", err)
			return
		}

		// Call the next handler if authentication passes
		next(w, r)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("User %s Hit endpoint", r.FormValue("user"))
		next(w, r)
	}
}

// list of all the middleware we have
var middleware = []func(http.HandlerFunc) http.HandlerFunc{
	authMiddleWare,
	loggingMiddleware,
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// if !authenticate(r) {
	// 	err := http.StatusUnauthorized
	// 	http.Error(w, "Invalid username or password", err)
	// 	return
	// }
	fmt.Fprintln(w, "Hello, welcome to my website!")
}

func main() {
	h := welcomeHandler
	for _, m := range middleware {
		h = m(h)
	}

	http.HandleFunc("/welcome", h)
	http.ListenAndServe(":8080", nil)
}
