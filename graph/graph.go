package graph

import (
	"net/http"

	"github.com/jonDufty/recipes/auth/rpc/authpb"
	"github.com/jonDufty/recipes/config"
	"github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
)

type App struct {
	Config  config.GraphConfig
	Clients *Clients
}

type Clients struct {
	Auth     authpb.Auth
	Cookbook cookbookpb.Cookbook
}

func New(cfg config.GraphConfig) *App {
	app := &App{
		Config: cfg,
		Clients: &Clients{
			Auth:     authpb.NewAuthProtobufClient(cfg.AuthEndpoint, &http.Client{}),
			Cookbook: cookbookpb.NewCookbookProtobufClient(cfg.CookbookEndpoint, &http.Client{}),
		},
	}

	return app
}
