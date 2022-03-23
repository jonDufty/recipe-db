package cookbook

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/config"
	"github.com/urfave/cli/v2"
)

type App struct {
	Config *config.CookbookConfig
	Ctx    context.Context
	DB     *sql.DB
}

func ServeHandler(c *cli.Context) error {
	cfg := config.NewCookbookConfig()

	// connect to database
	db, err := database.Connect(*cfg.DB)
	if err != nil {
		return errors.New("Could not connect to DBL " + err.Error())
	}

	fmt.Printf("Connected to DB at %s\n", cfg.DB.Address)

	app := App{
		Config: cfg,
		DB:     db,
	}

	bindAddr := fmt.Sprintf(":%d", app.Config.Port)
	log.Printf("Listening on %s", bindAddr)
	server := http.Server{
		Addr:         bindAddr,
		Handler:      chi.NewRouter(),
		IdleTimeout:  20 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	return server.ListenAndServe()
}
