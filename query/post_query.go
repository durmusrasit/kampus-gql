package query

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/durmusrasit/kampus-gql/loader"
	"github.com/durmusrasit/kampus-gql/models"
)

func (q *Query) CreatePost(ctx context.Context, args *struct {
	Title   string
	Url     string
	Content *string
	UserId  string
}) (*postResolver, error) {
	post := &models.Post{
		ID:      "1002",
		Title:   args.Title,
		Url:     args.Url,
		Content: *args.Content,
		Slug:    strings.ToLower(args.Title),
		UserID:  args.UserId,
	}

	fmt.Println("post", post)

	return &postResolver{post}, nil
}

func (q *Query) Posts(ctx context.Context) (*[]*postResolver, error) {
	posts, err := loader.ReadPosts()

	if err != nil {
		return nil, errors.New("An error occurred while reading posts" + err.Error())
	}

	var batch []*postResolver
	for _, p := range posts {
		post := postResolver{p}
		batch = append(batch, &post)
	}

	return &batch, nil
}

func (q *Query) Post(ctx context.Context, args struct {
	Slug string
	Id   string
}) (*postResolver, error) {
	post, err := q.GetPost(ctx, args.Slug, args.Id)
	if err != nil {
		return nil, err
	}

	return &postResolver{post}, nil
}

func (q *Query) GetPost(ctx context.Context, Slug string, Id string) (*models.Post, error) {
	posts, err := loader.ReadPosts()

	if err != nil {
		return nil, errors.New("An error occurred while reading posts" + err.Error())
	}

	for _, p := range posts {
		if p.ID == Id && p.Slug == Slug {
			return p, nil
		}
	}

	return nil, errors.New("Post not found")
}
