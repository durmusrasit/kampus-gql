package query

import pano_api "github.com/durmusrasit/kampus-gql/rpc/pano-api"

type Clients struct {
	PanoAPI pano_api.PanoAPI
}

type Query struct {
	Clients *Clients
}
