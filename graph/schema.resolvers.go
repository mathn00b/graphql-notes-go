package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mathnoob/graphql-notes-go/graph/generated"
	"github.com/mathnoob/graphql-notes-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCar(ctx context.Context) (*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "123",
			Text: "text",
			Done: false,
			User: nil,
		},
	}, nil
}

func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	return []*model.Car{
		{
			ID:    "123",
			Name:  "test-name",
			Color: "blue",
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
