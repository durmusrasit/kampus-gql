package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durmusrasit/kampus-gql/db/postgresql"
	"github.com/durmusrasit/kampus-gql/loader"
	"github.com/durmusrasit/kampus-gql/query"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	s, err := loader.ReadSchema("./schema.graphql")
	if err != nil {
		fmt.Println("An error occurred while reading schema. Error:", err)
		return
	}

	pgConfig := postgresql.PostgreSQLConfig{Host: "localhost", Port: 5432, Username: "postgres", Password: "yutubar123", DbName: "pano_db"}
	pgClient, pgError := postgresql.NewPostgreSQLConnect(pgConfig)
	if pgError != nil {
		fmt.Println("An error occurred while connect postgresql. Error:", pgError)
		return
	}

	db, err := loader.NewDB(pgClient)
	if err != nil {
		fmt.Println("An error occurred while load db. Error:", err)
	}

	schema := graphql.MustParseSchema(s, &query.Query{Db: db}, graphql.UseStringDescriptions())
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
