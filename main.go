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
					schema {
						query: Query
						mutation: Mutation
					}

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

					type Mutation {
						createPost(title: String!, url: String!, content: String, userId: String!): Post
					}
	`

	schema := graphql.MustParseSchema(s, &query.Query{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
