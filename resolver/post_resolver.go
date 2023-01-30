package resolver

import (
	pano_api "github.com/durmusrasit/kampus-gql/rpc/pano-api"
	"github.com/graph-gophers/graphql-go"
)

type postResolver struct {
	post *pano_api.Post
}

func (r *postResolver) ID() graphql.ID { return graphql.ID(r.post.Id) }

func (r *postResolver) Title() string { return r.post.Title }

func (r *postResolver) Url() string { return r.post.Url }

func (r *postResolver) Content() *string { return &r.post.Content }

func (r *postResolver) Slug() string { return r.post.Slug }

func (r *postResolver) UserID() string { return r.post.UserId }
