package resolver

import pano_api "github.com/durmusrasit/pano-api/rpc/pano-api"

type Clients struct {
	PanoAPI pano_api.PanoAPI
}

type Resolver struct {
	Clients *Clients
}
