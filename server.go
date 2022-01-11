package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/mathnoob/graphql-notes-go/graph/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mathnoob/graphql-notes-go/graph"
	"github.com/mathnoob/graphql-notes-go/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	u := NewUser()

	c := generated.Config{Resolvers: &graph.Resolver{
		AllCars:       []*model.Car{},
		CarsObservers: map[string]chan []*model.Car{},
	}}

	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
		if !u.getCurrentUser(ctx).HasRole(role) {
			// block calling the next resolver
			return nil, fmt.Errorf("Access denied")
		}

		// or let it pass through
		return next(ctx)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type User struct {
	ID   string
	Name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) getCurrentUser(ctx context.Context) *User {
	user, ok := ctx.Value("user").(*User)
	if !ok {
		return nil
	}

	u.ID = user.ID
	u.Name = user.Name

	return u
}

func (u *User) HasRole(role model.Role) bool {
	if u == nil {
		return false
	}

	return role.IsValid()
}
