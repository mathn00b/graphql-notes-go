package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/dgryski/trifles/uuid"
	"github.com/mathnoob/graphql-notes-go/graph/generated"
	"github.com/mathnoob/graphql-notes-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCar(ctx context.Context) (*model.Car, error) {
	car := &model.Car{
		ID:    uuid.UUIDv4(),
		Name:  gofakeit.Name(),
		Color: gofakeit.Color(),
	}

	r.AllCars = append(r.AllCars, car)

	r.mu.Lock()
	for _, observer := range r.CarsObservers {
		observer <- r.AllCars
	}
	r.mu.Unlock()

	return car, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Cars(ctx context.Context) (<-chan []*model.Car, error) {
	id := uuid.UUIDv4()
	cars := make(chan []*model.Car, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.CarsObservers, id)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	r.CarsObservers[id] = cars
	r.mu.Unlock()

	r.CarsObservers[id] <- r.AllCars
	return cars, nil
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
