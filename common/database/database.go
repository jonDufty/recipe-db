package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Address string `envconfig:"address"`
	User    string `envconfig:"user"`
	Pass    string `envconfig:"pass"`
	Name    string `envconfig:"name"`
}

func TestConnection(c Config) error {
	dsn := c.Address + c.Name
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
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

// Update

// Query

// Get

// Close
