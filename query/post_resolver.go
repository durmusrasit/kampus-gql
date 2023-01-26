package query

import (
	"github.com/durmusrasit/kampus-gql/models"
)

type postResolver struct {
	post *models.Post
}

//func (r *postResolver) ID() graphql.ID { return graphql.ID(r.post.ID) }

func (r *postResolver) ID() string { return r.post.ID.String() }

func (r *postResolver) Title() string { return r.post.Title }

func (r *postResolver) Url() string { return r.post.Url }

func (r *postResolver) Content() *string { return &r.post.Content }

func (r *postResolver) Slug() string { return r.post.Slug }

func (r *postResolver) UserID() string { return r.post.UserID }
