package cookbook

import (
	"context"

	rpc "github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
)

type Server struct {
	app *App
}

func NewServer(a *App) *Server {
	return &Server{a}
}

func (s *Server) GetRecipeById(context.Context, *rpc.GetRecipeByIdRequest) (*rpc.Recipe, error) {
	panic("implement me")
}
func (s *Server) InsertRecipe(context.Context, *rpc.Recipe) (*rpc.InsertRecipeResponse, error) {
	panic("implement me")
}
