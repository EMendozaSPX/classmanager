package main

import (
	"fmt"
	"github.com/emendoza/classmanager/pkg/Env"
)

func main() {
	/*
	// Creates the http route graphql
	http.HandleFunc("/graphql", Env.GraphQLHandler)

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

	*/

	fmt.Println(Env.GetSecretKey())
}
