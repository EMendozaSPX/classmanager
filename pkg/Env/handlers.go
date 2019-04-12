package Env

import (
	"encoding/json"
	"fmt"
	"github.com/emendoza/classmanager/pkg/Data"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"net/url"
)

type GraphqlQuery struct {
	Query string
	Variables map[string]interface{}
	OperationName string
}

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	// Apply standard http headers for sending and receiving information
	w.Header().Add("Content-Type", "application/json; charset=utf8")

	// get graphql query from http post request
	query := getQuery(r.URL.Query())

	if query != nil {
		// set graphql parameters
		params := graphql.Params{
			Schema: Data.Schema,
			RequestString: query.Query,
			VariableValues: query.Variables,
			OperationName: query.OperationName,
		}

		// produce a result
		result := graphql.Do(params)

		// write ok status code to http header
		w.WriteHeader(http.StatusOK)

		// format code as json
		buff, formatErr := json.MarshalIndent(result, "", "\t")
		if formatErr != nil {
			log.Println(formatErr)
		}

		// write to http response
		response, responseErr := w.Write(buff)
		if responseErr != nil {
			log.Println(responseErr)
		} else {
			fmt.Println("%d success", response)
		}
	}
}

// This function parses a http post response into a query object
func getQuery(d url.Values) *GraphqlQuery {
	// Graphql post responses have three main components, the query component contains the graphql query
	// https://graphql.org/learn/serving-over-http/
	query := d.Get("query")

	// Checks for query
	if query != "" {
		// create a variable named variables with a string array type
		variables := make(map[string]interface{}, len(d))

		// get the variables in the form of a json object from the post request
		variablesJson := d.Get("variables")

		// serialize the json object into the variables variable
		if err := json.Unmarshal([]byte(variablesJson), &variables); err != nil {
			log.Println(err)
		}

		// return the query
		return &GraphqlQuery{
			Query: query,
			Variables: variables,
			OperationName: d.Get("operationName"),
		}
	}

	// return nil if the variables are empty
	return nil
}