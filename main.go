package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durmusrasit/kampus-gql/resolver"
	"github.com/durmusrasit/kampus-gql/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	pano_api "github.com/kamp-us/pano-api/rpc/pano-api"
)

func main() {

	s, err := schema.ReadSchema("./schema/schema.graphql")
	if err != nil {
		fmt.Println("An error occurred while reading schema. Error:", err)
		return
	}

	panoapiClient := pano_api.NewPanoAPIProtobufClient("http://localhost:8080", &http.Client{})

	clients := resolver.Clients{
		PanoAPI: panoapiClient,
	}

	schema := graphql.MustParseSchema(s, &resolver.Resolver{Clients: &clients}, graphql.UseStringDescriptions())
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("Listening to :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
