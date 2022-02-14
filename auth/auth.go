package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonDufty/recipes/common"
	"github.com/jonDufty/recipes/config"
	"github.com/urfave/cli/v2"
)

type App struct {
	Config *config.AuthConfig
	Router *chi.Mux
}

func ServeHandler(context *cli.Context) error {
	config := config.NewAuthConfig()

	app := App{}
	app.Config = config

	r := common.NewRouter()
	app.Router = r

	// Add routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🔐")
	})

	log.Printf("Listening on port 80 ....")
	return http.ListenAndServe(":80", r)

}
