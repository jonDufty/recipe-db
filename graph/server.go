package graph

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jonDufty/recipes/config"
	"github.com/jonDufty/recipes/graph/generated"
	"github.com/urfave/cli/v2"
)

func ServeGraph(c *cli.Context) error {

	cfg := config.NewGraphConfig()
	app := New(*cfg)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{app},
			Directives: generated.DirectiveRoot{
				IsAuthenticated: isAuthenticatedDirective,
			},
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}

func isAuthenticatedDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return next(ctx)
}
