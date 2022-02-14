package auth

import (
	"context"

	pb "github.com/jonDufty/recipes/auth/rpc/authpb"
)

type Server struct {
	app *App
}

func NewServer(a *App) *Server {
	return &Server{a}
}

func (s *Server) SayHello(c context.Context, r *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {

	return &pb.SayHelloResponse{
		Greeting: "Hello " + r.Name,
	}, nil
}
