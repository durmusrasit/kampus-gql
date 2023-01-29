package panoapi

import pano_api "github.com/durmusrasit/kampus-gql/rpc/pano-api"

type PanoAPIProtobufClient struct {
	client pano_api.PanoAPI
}

type PanoAPI struct {
	PanoAPI *PanoAPIProtobufClient
}

func NewPanoAPI(pano_api *PanoAPIProtobufClient) (*PanoAPI, error) {
	return &PanoAPI{pano_api}, nil
}
