package main

import (
	"fmt"
	"github.com/emendoza/classmanager/pkg/Data"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	// h is a callback function that handles http requests and responses
	// h handles all graphql data requests and responses in the route /graphql
	h := handler.New(&handler.Config{
		Schema: &Data.Schema,
		Pretty: true,
		GraphiQL: false,
		Playground: true,
	})

	// Creates the http route graphql
	http.Handle("/graphql", h)

	// static file handler, a callback function that serves static (React) files
	static := http.FileServer(http.Dir("web"))

	// print instructions to console
	fmt.Println("open localhost:3000 in web browser")

	// serve static files at http root
	http.Handle("/", static)

	// deploy dev server
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
