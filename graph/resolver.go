package graph

import (
	"github.com/mathnoob/graphql-notes-go/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// All cars since launching the GraphQL endpoint
	AllCars []*model.Car
	// All active subscriptions
	CarsObservers map[string]chan []*model.Car
	mu            sync.Mutex
}
