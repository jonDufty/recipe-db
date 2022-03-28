package mappers

import (
	"time"

	"github.com/jonDufty/recipes/auth/models/user"
	"github.com/jonDufty/recipes/auth/rpc/authpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func UserFromProto(u *authpb.User, hash string) *user.User {

	return &user.User{
		FullName:     u.FullName,
		Email:        u.Email,
		TimeCreated:  time.Now(),
		PasswordHash: hash,
	}
}

func UserFromSignupRequst(r *authpb.SignupRequest, hash string) *user.User {

	return &user.User{
		FullName:     r.FullName,
		Email:        r.Email,
		TimeCreated:  time.Now(),
		PasswordHash: hash,
	}
}

func ProtoFromUser(u *user.User) *authpb.User {
	return &authpb.User{
		UserId:    int64(u.ID),
		FullName:  u.FullName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.TimeCreated),
	}
}
