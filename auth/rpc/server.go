package auth

import (
	"context"

	"github.com/jonDufty/recipes/auth"
	pb "github.com/jonDufty/recipes/auth/rpc/authpb"
)

type Server struct {
	app *auth.App
}

func NewServer(a *auth.App) *Server {
	return &Server{a}
}

func (s *Server) SayHello(c context.Context, r *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {

	return &pb.SayHelloResponse{
		Greeting: "Hello" + r.Name,
	}, nil
}
