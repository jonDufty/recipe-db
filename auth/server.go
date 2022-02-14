package auth

import (
	"context"

	pb "github.com/jonDufty/recipes/auth/rpc/authpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	app *App
}

var fakeUser = pb.User{
	UserId:    1234,
	FullName:  "Test User",
	Email:     "test@example.com",
	CreatedAt: timestamppb.Now(),
}

func NewServer(a *App) *Server {
	return &Server{a}
}

func (s *Server) SayHello(c context.Context, r *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {

	return &pb.SayHelloResponse{
		Greeting: "Hello " + r.Name,
	}, nil
}

func (s *Server) GetUserByEmail(c context.Context, r *pb.ByEmailRequest) (*pb.User, error) {
	return &fakeUser, nil
}

func (s *Server) GetUserById(c context.Context, r *pb.ByIdRequest) (*pb.User, error) {
	return &fakeUser, nil
}
