package main

import (
	"fmt"
	"github.com/emendoza/classmanager/pkg/Schema"
	"github.com/emendoza/classmanager/pkg/GraphqlHandler"
	"log"
	"net/http"
)

func main() {
	// Create a handler for the graphql queries
	h := GraphqlHandler.New(&GraphqlHandler.Config{
		Schema: &Schema.Schema,
	})

	// Creates the http route graphql
	http.Handle("/graphql", h)

	// serve index.html as a static file
	root := http.FileServer(http.Dir("./web/build"))

	// serve index.html at http root
	http.Handle("/", root)

	// print instructions to console
	fmt.Println("open http://localhost:3030 in web browser")

	// deploy dev server
	err := http.ListenAndServe(":3030", nil)

	if err != nil {
		log.Fatal(err)
	}
}
