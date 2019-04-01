package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
	// h is a callback function that handles http requests and responses
	// h handles all graphql data requests and responses in the route /graphql
	h := handler.New(&handler.Config{

	})

	// Creates the http route graphql
	http.Handle("/graphql", h)
	 */

	// static file handler, a callback function that serves static (React) files
	static := http.FileServer(http.Dir("static"))

	// print instructions to console
	fmt.Println("open localhost:8000 in web browser")

	// serve static files at http root
	http.Handle("/", static)

	// deploy dev server
	http.ListenAndServe(":8000", nil)
}
