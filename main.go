package main

import (
	"log"
	"net/http"

	"github.com/durmusrasit/kampus-gql/query"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	s := `
					type Post {
						ID: String!
						Title: String!
						Url: String!
						Content: String
						Slug: String!
						UserID: String!
					}

					type Query {
						post(slug: String!, id: String!): Post!
						posts(): [Post!]
					}
	`

	schema := graphql.MustParseSchema(s, &query.Query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
