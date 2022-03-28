package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func newTestDBConfig(testName string) Config {

	// Check if we're in docker or not
	addr := "mysql:3306"
	if _, err := os.Stat("/.dockerenv"); err != nil {
		addr = "127.0.0.1:45000"
	}

	return Config{
		Address: fmt.Sprintf("root@tcp(%s)/", addr),
		User:    "root",
		Name:    "test_recipe" + testName,
	}
}

func mustExec(db *sql.Tx, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func initTestDB(db *sql.DB, cfg Config) error {

	fmt.Println("Creating DB " + cfg.Name)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	mustExec(tx, "DROP DATABASE IF EXISTS "+cfg.Name)
	mustExec(tx, "CREATE DATABASE "+cfg.Name)
	mustExec(tx, "USE "+cfg.Name)

	fmt.Println("Migrating up")
	handler, err := ConnectDB(fmt.Sprintf("mysql://%s/%s", cfg.Address, cfg.Name))
	if err != nil {
		return err
	}

	defer handler.Close()

	return handler.Up(context.Background())
}

func NewTestDBConnection() (*sql.DB, func()) {

	cfg := newTestDBConfig("basic_test")
	db, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	initTestDB(db, cfg)

	closer := func() {
		_, err := db.Exec("DROP DATABASE IF EXISTS " + cfg.Name)
		if err != nil {
			panic(err)
		}
		db.Close()
	}

	return db, closer
}
