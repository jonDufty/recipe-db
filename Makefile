PREFIX=github.com/jonDufty/recipes

all: twirp-auth

twirp-auth:
	protoc \
	--twirp_out=. \
	--go_out=. \
	auth/rpc/proto/auth.proto

clean:
	rm -rf auth/rpc/authpb/*