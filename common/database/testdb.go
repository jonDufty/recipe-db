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
		fmt.Println("Not in docker, using localhost")
		addr = "127.0.0.1:45000"
	}

	return Config{
		Address: fmt.Sprintf("root@tcp(%s)/", addr),
		User:    "root",
		Name:    testName,
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

	for _, query := range Schema {
		mustExec(tx, query.Sql)
	}

	return nil
}

func NewTestDBConnection() (*sql.DB, func()) {

	cfg := newTestDBConfig("")
	db, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	cfg.Name = "test_recipes"
	closer := func() {
		_, err := db.Exec("DROP DATABASE IF EXISTS " + cfg.Name)
		if err != nil {
			panic(err)
		}
		db.Close()
	}

	if err := initTestDB(db, cfg); err != nil {
		closer()
		panic(err)
	}

	// Close then reconnect with test db
	db.Close()
	db, err = Connect(cfg)
	if err != nil {
		panic(err)
	}

	return db, closer
}
