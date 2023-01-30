package query

import (
	"context"

	pano_api "github.com/durmusrasit/kampus-gql/rpc/pano-api"
)

type CreatePostInput struct {
	Title   string
	Url     string
	Content *string
	UserId  string
}

type CreatePostArgs struct {
	Input *CreatePostInput
}

func (q *Query) CreatePost(ctx context.Context, args *CreatePostArgs) (*postResolver, error) {
	createdPost, err := q.Clients.PanoAPI.CreatePost(ctx, &pano_api.CreatePostRequest{Title: args.Input.Title, Url: args.Input.Url, Content: *args.Input.Content, UserId: args.Input.UserId})

	if err != nil {
		return nil, err
	}

	return &postResolver{createdPost.Post}, nil
}

type GetPostInput struct {
	Slug string
	Id   string
}

type GetPostArgs struct {
	Input *GetPostInput
}

func (q *Query) Post(ctx context.Context, args *GetPostArgs) (*postResolver, error) {
	response, err := q.Clients.PanoAPI.GetBatchPosts(ctx, &pano_api.GetBatchPostsRequest{Ids: []string{args.Input.Id}})

	if err != nil {
		return nil, err
	}

	return &postResolver{response.Posts[0]}, nil
}

func (q *Query) Posts(ctx context.Context) (*[]*postResolver, error) {
	response, err := q.Clients.PanoAPI.GetPosts(ctx, &pano_api.GetPostsRequest{})

	if err != nil {
		return nil, err
	}

	var batch []*postResolver
	for _, p := range response.Posts {
		post := postResolver{p}
		batch = append(batch, &post)
	}

	return &batch, nil
}
