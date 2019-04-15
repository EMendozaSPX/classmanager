package main

import (
	"fmt"
	"github.com/emendoza/classmanager/pkg/Data"
	"github.com/emendoza/classmanager/pkg/GraphqlHandler"
	"log"
	"net/http"
)

func main() {
	// Create a handler for the graphql queries
	h := GraphqlHandler.New(&GraphqlHandler.Config{
		Schema: &Data.Schema,
	})

	// Creates the http route graphql
	http.Handle("/graphql", h)

	// static file handler, a callback function that serves static (React) files
	static := http.FileServer(http.Dir("web"))

	// print instructions to console
	fmt.Println("open localhost:3030 in web browser")

	// serve static files at http root
	http.Handle("/", static)

	// deploy dev server
	err := http.ListenAndServe(":3030", nil)

	if err != nil {
		log.Fatal(err)
	}
}
