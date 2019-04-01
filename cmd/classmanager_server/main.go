package main

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	// h is a callback function that handles http requests and responses
	// h handles all graphql data requests and responses in the route /graphql
	h := handler.New(&handler.Config{

	})

	// Creates the http route graphql
	http.Handle("/graphql", h)

	// deploy dev server
	http.ListenAndServe(":8000", nil)
}
