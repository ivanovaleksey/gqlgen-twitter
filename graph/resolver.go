package graph

import (
	"context"
	"github.com/ivanovaleksey/twitter/graph/model"
	"github.com/ivanovaleksey/twitter/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	storage Storage
}

func NewResolver() *Resolver {
	return &Resolver{storage: storage.NewInMemory()}
}

type Storage interface {
	CreatePost(ctx context.Context, newPost model.NewPost) (model.Post, error)
	GetLatestPostsByUser(ctx context.Context, params storage.GetLatestPostsByUserParams) ([]model.Post, int, error)
}
