package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jonDufty/recipes/auth/routes"
	"github.com/jonDufty/recipes/config"
	"github.com/urfave/cli/v2"
)

type App struct {
	Config *config.AuthConfig
}

func ServeHandler(context *cli.Context) error {
	config := config.NewAuthConfig()

	app := App{
		Config: config,
	}

	bindAddr := fmt.Sprintf(":%d", app.Config.Port)
	log.Printf("Listening on %s", bindAddr)
	server := http.Server{
		Addr:         bindAddr,
		Handler:      routes.NewRouter(),
		IdleTimeout:  20 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	return server.ListenAndServe()

}
