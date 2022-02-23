package database

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/russross/meddler"
)

type Config struct {
	Address string `envconfig:"address"`
	User    string `envconfig:"user"`
	Pass    string `envconfig:"pass"`
	Name    string `envconfig:"name"`
}

const DBKey = "database"

func Get(ctx context.Context) (meddler.DB, error) {
	if db, ok := ctx.Value(DBKey).(meddler.DB); ok {
		return db, nil
	}
	return nil, errors.New("DB not in context")
}

func dbContext(ctx context.Context, db meddler.DB) context.Context {
	return context.WithValue(ctx, DBKey, db)
}

// Connect to database
func Connect(c Config) (*sql.DB, error) {
	dsn := c.Address + c.Name
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Middleware
func Middleware(db meddler.DB) func(handler http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		f := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(dbContext(r.Context(), db)))
		}
		return http.HandlerFunc(f)
	}
}

// Insert
func Insert(ctx context.Context, table string, src interface{}) error {
	db, err := Get(ctx)
	if err != nil {
		return err
	}

	tx, err := db.(*sql.DB).BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil
	}

	err = meddler.Insert(tx, table, src)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		err = tx.Commit()
	}

	return err
}

// Update

// Query

// Get

// Close
