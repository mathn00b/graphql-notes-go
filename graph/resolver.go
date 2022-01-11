package graph

import (
	"github.com/brianvoe/gofakeit"
	"github.com/dgryski/trifles/uuid"
	"github.com/mathnoob/graphql-notes-go/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Root *Root
	// All active subscriptions
	CarsObservers map[string]chan []*model.Car
	mu            sync.Mutex
}

type Root struct {
	Cars  []*model.Car
	Todos []*model.Todo
}

func NewCarsStub(count int) []*model.Car {
	cars := make([]*model.Car, count)

	for i := 0; i < count; i++ {
		cars[i] = newCar()
	}

	return cars
}

func newCar() *model.Car {
	return &model.Car{
		ID:    uuid.UUIDv4(),
		Name:  gofakeit.Name(),
		Color: gofakeit.Color(),
	}
}

func NewTodosStub(count int) []*model.Todo {
	todos := make([]*model.Todo, count)

	for i := 0; i < count; i++ {
		todos[i] = newTodo()
	}

	return todos
}

func newTodo() *model.Todo {
	return &model.Todo{
		ID:   uuid.UUIDv4(),
		Text: gofakeit.Country(),
		Done: gofakeit.Bool(),
		User: newUser(),
	}
}

func newUser() *model.User {
	return &model.User{
		ID:   uuid.UUIDv4(),
		Name: gofakeit.Name(),
	}
}
