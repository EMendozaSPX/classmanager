package main

import (
	"fmt"
	"github.com/emendoza/classmanager/pkg/Env"
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
	fmt.Printf("open http://localhost%s in web browser \n", Env.GetPort())

	// deploy dev server
	err := http.ListenAndServe(Env.GetPort(), nil)

	if err != nil {
		log.Fatal(err)
	}
}
