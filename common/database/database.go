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

type DBkey string

const dBKey DBkey = "database"

func Get(ctx context.Context) (meddler.DB, error) {
	if db, ok := ctx.Value(dBKey).(meddler.DB); ok {
		return db, nil
	}
	return nil, errors.New("DB not in context")
}

func DbContext(ctx context.Context, db meddler.DB) context.Context {
	return context.WithValue(ctx, dBKey, db)
}

// Connect to database
func Connect(c Config) (*sql.DB, error) {
	dsn := c.Address + c.Name + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	return db, nil
}

// Middleware
func Middleware(db meddler.DB) func(handler http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		f := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(DbContext(r.Context(), db)))
		}
		return http.HandlerFunc(f)
	}
}

func StartTx(ctx context.Context) (*sql.Tx, error) {
	db, err := Get(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := db.(*sql.DB).BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Insert
func Insert(ctx context.Context, table string, src interface{}) error {

	tx, err := StartTx(ctx)
	if err != nil {
		return err
	}

	err = meddler.Insert(tx, table, src)
	if err != nil {
		e := tx.Rollback()
		if e != nil {
			return errors.New(err.Error() + e.Error())
		}
		return err
	} else {
		err = tx.Commit()
	}

	return err
}

// Update
func Update(ctx context.Context, table string, src interface{}) error {
	tx, err := StartTx(ctx)
	if err != nil {
		return err
	}

	err = meddler.Insert(tx, table, src)
	if err != nil {
		e := tx.Rollback()
		if e != nil {
			return errors.New(e.Error() + e.Error())
		}
		return err
	} else {
		err = tx.Commit()
	}

	return err
}

// Query

// Get - Users meddler.QueryRow to get a single row result

// Close
