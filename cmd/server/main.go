package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println(e)
	/*
	// h is a callback function that handles http requests and responses
	// h handles all graphql data requests and responses in the route /graphql
	h := handler.New(&handler.Config{

	})

	// Creates the http route graphql
	http.Handle("/graphql", h)
	 */

	// static file handler, a callback function that serves static (React) files
	static := http.FileServer(http.Dir("web"))

	// print instructions to console
	fmt.Println("open localhost:3000 in web browser")

	// serve static files at http root
	http.Handle("/", static)

	// deploy dev server
	http.ListenAndServe(":3000", nil)
}
