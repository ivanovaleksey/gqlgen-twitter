package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ivanovaleksey/gqlgen-twitter/graph/generated"
	"github.com/ivanovaleksey/gqlgen-twitter/graph/model"
	"github.com/ivanovaleksey/gqlgen-twitter/storage"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post, err := r.storage.CreatePost(ctx, input)
	if err != nil {
		return nil, err
	}

	r.pubSub.Pub(ctx, post)

	return &post, nil
}

func (r *queryResolver) LatestUserPosts(ctx context.Context, userID int64, limit *int64, offset *int64) (*model.PostsList, error) {
	params := storage.GetLatestPostsByUserParams{
		UserID: userID,
	}
	if limit != nil {
		params.Limit = int(*limit)
	}
	if offset != nil {
		params.Offset = int(*offset)
	}
	items, total, err := r.storage.GetLatestPostsByUser(ctx, params)
	if err != nil {
		return nil, err
	}
	itemsPtr := make([]*model.Post, 0, len(items))
	for _, item := range items {
		itemsPtr = append(itemsPtr, &item)
	}
	return &model.PostsList{
		Items: itemsPtr,
		Total: int64(total),
	}, nil
}

func (r *subscriptionResolver) NewPostEvent(ctx context.Context, userID int64) (<-chan *model.Post, error) {
	sub, err := r.pubSub.Sub(ctx, userID)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		r.pubSub.Unsub(ctx, sub.ID, userID)
	}()

	return sub.Posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
