package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durmusrasit/kampus-gql/query"
	"github.com/durmusrasit/kampus-gql/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	pano_api "github.com/durmusrasit/kampus-gql/rpc/pano-api"
)

func main() {

	s, err := schema.ReadSchema("./schema/schema.graphql")
	if err != nil {
		fmt.Println("An error occurred while reading schema. Error:", err)
		return
	}

	panoapiClient := pano_api.NewPanoAPIProtobufClient("http://localhost:8080", &http.Client{})

	clients := query.Clients{
		PanoAPI: panoapiClient,
	}

	schema := graphql.MustParseSchema(s, &query.Query{Clients: &clients}, graphql.UseStringDescriptions())
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("Listening to :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
