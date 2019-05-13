package GraphqlHandler

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// This graphql handler is the work of the graphql-go team, https://github.com/graphql-go/handler/blob/master/handler.go
// The work has been modified under the terms of the MIT License, in order to suit this particular application of graphql

// Declare constant strings for determining the http post format
const (
	ContentTypeJson    =        "application/json"
	ContentTypeGraphQL =        "application/graphql"
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

type ResultCallbackFn func(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte)

type RootObjectFn func(ctx context.Context, r *http.Request) map[string]interface{}

// Create a http handler struct
type Handler struct {
	Schema           *graphql.Schema
	rootObjectFn     RootObjectFn
	resultCallbackFn ResultCallbackFn
	formatErrorFn      func(err error) gqlerrors.FormattedError
}

// This struct stores the parameters of the graphql query
type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

// This struct is the same as the previous with a slight change in typing in order to avoid serialization errors
type requestOptionsCompatability struct {
	Query         string `json:"query" url:"query" schema:"query"`
	Variables     string `json:"variables" url:"variables" schema:"variables"`
	OperationName string `json:"operationName" url:"operationName" schema:"operationName"`
}

// This function formats a graphql query using the variables from the post request
func getFromForm(values url.Values) *RequestOptions {
	// Gets the grahphql query from the post request
	query := values.Get("query")
	if query != "" {
		variables := make(map[string]interface{}, len(values))
		variablesStr := values.Get("variables")
		if err := json.Unmarshal([]byte(variablesStr), &variables); err != nil {
			log.Println(err)
		}

		return &RequestOptions{
			Query:         query,
			Variables:     variables,
			OperationName: values.Get("operationName"),
		}
	}

	return nil
}

func NewRequestOptions(r *http.Request) *RequestOptions {
	if reqOpt := getFromForm(r.URL.Query()); reqOpt != nil {
		return reqOpt
	}

	if r.Method != http.MethodPost {
		return &RequestOptions{}
	}

	if r.Body == nil {
		return &RequestOptions{}
	}

	contentTypeStr := r.Header.Get("Content-Type")
	contentTypeTokens := strings.Split(contentTypeStr, ";")
	contentType := contentTypeTokens[0]

	switch contentType {
	case ContentTypeGraphQL:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return &RequestOptions{}
		}
		return &RequestOptions{
			Query: string(body),
		}

	case ContentTypeFormURLEncoded:
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return &RequestOptions{}
		}

		if reqOpt := getFromForm(r.PostForm); reqOpt != nil {
			return reqOpt
		}
		return &RequestOptions{}

	case ContentTypeJson:
		fallthrough

	default:
		var opts RequestOptions
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return &opts
		}
		if err := json.Unmarshal(body, &opts); err != nil {
			log.Println(err)
			var optsCompatible requestOptionsCompatability
			if err := json.Unmarshal(body, &optsCompatible); err != nil {
				log.Println(err)
			}
			if err := json.Unmarshal([]byte(optsCompatible.Variables), &opts.Variables); err != nil {
				log.Println(err)
			}
		}
		return &opts
	}
}

func (h *Handler) ContextHandler(contextValue context.Context, w http.ResponseWriter, r *http.Request) {
	opts := NewRequestOptions(r)

	tokenStr := r.Header.Get("authorization")

	var token string

	if tokenStr == "" {
		token = ""
	} else {
		token = strings.Split(tokenStr, "Bearer ")[1]
	}

	ctx := context.WithValue(contextValue, "token", token)

	params := graphql.Params{
		Schema: *h.Schema,
		RequestString: opts.Query,
		VariableValues: opts.Variables,
		OperationName: opts.OperationName,
		Context: ctx,
	}
	if h.rootObjectFn != nil {
		params.RootObject = h.rootObjectFn(ctx, r)
	}
	result := graphql.Do(params)

	if formatErrorFn := h.formatErrorFn; formatErrorFn != nil && len(result.Errors) > 0 {
		formatted := make([]gqlerrors.FormattedError, len(result.Errors))
		for i, formattedError := range result.Errors {
			formatted[i] = formatErrorFn(formattedError.OriginalError())
		}
		result.Errors = formatted
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	var buff []byte
	w.WriteHeader(http.StatusOK)
	buff, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Println(err)
	}

	if _, err := w.Write(buff); err != nil {
		log.Println(err)
	}

	if h.resultCallbackFn != nil {
		h.resultCallbackFn(ctx, &params, result, buff)
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ContextHandler(r.Context(), w, r)
}

type Config struct {
	Schema           *graphql.Schema
	RootObjectFn     RootObjectFn
	ResultCallbackFn ResultCallbackFn
	FormatErrorFn    func(err error) gqlerrors.FormattedError
}

func New(p *Config) *Handler {
	if p == nil {
		log.Panic("Undefined config")
	}
	if p.Schema == nil {
		log.Panic("Undefined Schema")
	}
	return &Handler{
		Schema:           p.Schema,
		rootObjectFn:     p.RootObjectFn,
		resultCallbackFn: p.ResultCallbackFn,
		formatErrorFn:    p.FormatErrorFn,
	}
}