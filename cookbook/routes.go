package cookbook

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	db "github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
)

func NewRouter(a *App) http.Handler {
	r := chi.NewRouter()

	// Add common middleware
	r.Use(middleware.Logger)
	r.Use(db.Middleware(a.DB))

	r.Get("/system/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🔐")
	})

	// Add twirp routes
	r.Group((func(chi.Router) {
		// Add middleware
		twirpServer := cookbookpb.NewCookbookServer(NewServer(a))
		r.Mount("/cookbook/twirp/", http.StripPrefix("/cookbook", twirpServer))
	}))

	return r
}
