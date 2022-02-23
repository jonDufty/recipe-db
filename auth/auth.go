package auth

import (
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
	DB     *sql.DB
}

func ServeHandler(context *cli.Context) error {
	c := config.NewAuthConfig()

	// connect to database
	db, err := database.Connect(*c.DB)
	if err != nil {
		return errors.New("Could not connect to DB: " + err.Error())
	}

	fmt.Printf("Connected to DB at %s\n", c.DB.Address)

	app := App{
		Config: c,
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
