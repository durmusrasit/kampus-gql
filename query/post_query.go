package query

import (
	"context"
	"errors"
	"strings"

	"github.com/durmusrasit/kampus-gql/models"
)

type PostInput struct {
	Title   string
	Url     string
	Content *string
	UserId  string
}

func (q *Query) CreatePost(ctx context.Context, args *struct {
	Input *PostInput
}) (*postResolver, error) {

	post := &models.Post{
		Title:   args.Input.Title,
		Url:     args.Input.Url,
		Content: *args.Input.Content,
		Slug:    strings.ToLower(args.Input.Title),
		UserID:  args.Input.UserId,
	}

	result := q.Db.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &postResolver{post}, nil
}

func (q *Query) Posts(ctx context.Context) (*[]*postResolver, error) {
	posts, err := q.Db.ReadPosts()

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
	posts, err := q.Db.ReadPosts()

	if err != nil {
		return nil, errors.New("An error occurred while reading posts" + err.Error())
	}

	for _, p := range posts {
		if p.ID.String() == Id && p.Slug == Slug {
			return p, nil
		}
	}

	return nil, errors.New("Post not found")
}
