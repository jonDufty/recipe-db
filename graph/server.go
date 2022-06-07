package graph

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/jonDufty/recipes/common/middleware"
	"github.com/jonDufty/recipes/config"
	"github.com/jonDufty/recipes/graph/generated"
	"github.com/urfave/cli/v2"
)

func ServeGraph(c *cli.Context) error {

	cfg := config.NewGraphConfig()
	app := New(*cfg)

	routes := chi.NewRouter()
	routes.Use(middleware.AuthMiddleware(app.Clients.Auth))

	gql := gqlHandler(app)

	routes.Get("/system/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "graph OK")
	})

	routes.Handle("/", playground.Handler("GraphQL playground", "/query"))
	routes.Handle("/query", gql)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}

func gqlHandler(app *App) http.Handler {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{app},
			Directives: generated.DirectiveRoot{
				IsAuthenticated: isAuthenticatedDirective,
			},
		},
	))

	return srv
}

func isAuthenticatedDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if !middleware.IsAuthFromContext(ctx) {
		return nil, errors.New("unauthenticated")
	}
	return next(ctx)
}
