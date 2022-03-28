package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/config"
	"github.com/urfave/cli/v2"
)

type App struct {
	Config *config.AuthConfig
	Ctx    context.Context
	DB     *sql.DB
}

func ServeHandler(c *cli.Context) error {
	cfg := config.NewAuthConfig()

	// connect to database
	db, err := database.Connect(*cfg.DB)
	if err != nil {
		return errors.New("Could not connect to DB: " + err.Error())
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
		Handler:      NewRouter(&app),
		IdleTimeout:  20 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	return server.ListenAndServe()

}

func NewTestApp(c *config.AuthConfig) *App {

	app := App{
		Config: c,
		Ctx:    context.Background(),
	}

	return &app
}
