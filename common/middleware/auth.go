package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jonDufty/recipes/auth/rpc/authpb"
)

const cookieKey = "sesh"
const authenticatedKey = "authenticated"

func AuthMiddleware(client authpb.Auth) func(handler http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {

			if r.Context().Value(authenticatedKey).(bool) {
				next.ServeHTTP(w, r)
			}

			cookie, err := r.Cookie(cookieKey)
			if err != nil || cookie.Expires.Before(time.Now()) {
				log.Println("Invalid or missing cookie")
				r = r.WithContext(context.Background())
			}

			token := cookie.Value
			req := authpb.ValidateTokenRequest{
				Token: token,
			}
			resp, err := client.ValidateToken(r.Context(), &req)
			if err == nil && resp.Valid {
				r = r.WithContext(context.WithValue(r.Context(), authenticatedKey, true))
			}

			next.ServeHTTP(w, r)

		}
		return http.HandlerFunc(f)
	}
}

func IsAuthFromContext(ctx context.Context) bool {
	return ctx.Value(authenticatedKey).(bool)
}
