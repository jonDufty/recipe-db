package auth

import (
	"context"
	"log"

	"github.com/jonDufty/recipes/auth/mappers"
	"github.com/jonDufty/recipes/auth/models/user"
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

	if r.Email == "" || r.Password == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "Some fields are empty")
	}

	u, err := user.GetUserByEmail(c, r.Email)
	if err != nil {
		log.Printf("User of email: %s not found", r.Email)
		log.Print(err.Error())
		return nil, twirp.NewError(twirp.InvalidArgument, "")
	}

	err = crypto.CheckPassword(r.Password, u.PasswordHash)
	if err != nil {
		log.Printf("Password for email: %s does not match hash\nExpected: %s\n Received %s", r.Email, u.PasswordHash, r.Password)
		log.Print(err.Error())
		return nil, twirp.NewError(twirp.InvalidArgument, "")
	}

	token, err := GenerateToken(u)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, "Failed to generate token")
	}

	return &pb.LoginResponse{
		Token: token,
		User:  mappers.ProtoFromUser(u),
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

	token, err := GenerateToken(requestUser)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, "Failed to generate token")
	}

	return &pb.LoginResponse{
		Token: token,
		User:  mappers.ProtoFromUser(requestUser),
	}, nil
}

func (s *Server) ValidateToken(c context.Context, r *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	_, err := ValidateToken(r.Token)
	if err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, "Invalid token")
	}

	return &pb.ValidateTokenResponse{
		Valid: true,
	}, nil
}
