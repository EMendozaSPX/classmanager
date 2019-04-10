package Env

import (
	"encoding/json"
	"fmt"
	"github.com/emendoza/classmanager/pkg/Data"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf8")
	query := r.URL.Query().Get("query")
	variables := make(map[string]interface{}, len(r.URL.Query()))
	variablesStr := r.URL.Query().Get("variables")
	if err := json.Unmarshal([]byte(variablesStr), &variables); err != nil {
		log.Println(err)
	}

	params := graphql.Params{
		Schema: Data.Schema,
		RequestString: query,
		VariableValues: variables,
		OperationName: r.URL.Query().Get("operationName"),
	}

	result := graphql.Do(params)

	w.WriteHeader(http.StatusOK)
	buff, formatErr := json.MarshalIndent(result, "", "\t")
	if formatErr != nil {
		log.Println(formatErr)
	}
	response, responseErr := w.Write(buff)
	if responseErr != nil {
		log.Println(responseErr)
	} else {
		fmt.Println("%d success", response)
	}
}