package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durmusrasit/kampus-gql/query"
	"github.com/durmusrasit/kampus-gql/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	s, err := schema.ReadSchema("./schema.graphql")
	if err != nil {
		fmt.Println("An error occurred while reading schema. Error:", err)
		return
	}

	//panoapiClient := pano_api.NewPanoAPIProtobufClient("http://localhost:8080", &http.Client{})

	schema := graphql.MustParseSchema(s, &query.Query{ /*PanoAPI: panoapiClient*/ }, graphql.UseStringDescriptions())
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
