package auth

import (
	"context"
	"log"

	"github.com/jonDufty/recipes/auth/mappers"
	pb "github.com/jonDufty/recipes/auth/rpc/authpb"
	"github.com/jonDufty/recipes/common/crypto"
	"github.com/twitchtv/twirp"
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

func (s *Server) LoginWithPassword(c context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Token: "fake-token",
	}, nil
}

func (s *Server) SignupWithPassword(c context.Context, r *pb.SignupRequest) (*pb.LoginResponse, error) {

	if r.FullName == "" || r.Email == "" || r.Password == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "Some fields are empty")
	}

	hash, err := crypto.HashPassword(r.Password)
	if err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, "Failed to hash password")
	}

	requestUser := mappers.UserFromSignupRequst(r, hash)

	err = requestUser.InsertUser(c)
	if err != nil {
		twirp.NewError(twirp.Internal, "")
		log.Printf("Error inserting user into db: %v", requestUser)
	}

	log.Print("User created successfully")

	return &pb.LoginResponse{
		Token: "fake-token",
	}, nil
}

func (s *Server) ValidateToken(c context.Context, r *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return &pb.ValidateTokenResponse{
		Valid: true,
	}, nil
}
