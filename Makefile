PREFIX=github.com/jonDufty/recipes

all: twirp-auth

twirp-auth:
	protoc \
	--twirp_out=. \
	--go_out=. \
	auth/rpc/proto/auth.proto

twirp-cookbook:
	protoc \
	--twirp_out=. \
	--go_out=. \
	cookbook/rpc/proto/cookbook.proto

clean:
	rm -rf auth/rpc/authpb/*
	rm -rf auth/rpc/cookbookpb/*
