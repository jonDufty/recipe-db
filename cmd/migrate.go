package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/db-journey/migrate/v2"
	"github.com/db-journey/migrate/v2/file"
	mysqlMigrationDriver "github.com/db-journey/mysql-driver"
	cli "github.com/urfave/cli/v2"
)

func MigrateUp(ctx *cli.Context) error {
	fmt.Println("Migrating up ...")

	dsn := fmt.Sprint("mysql://root@(mysql:3306)/recipedb")

	handler, err := ConnectDB(dsn)
	if err != nil {
		return err
	}

	defer handler.Close()

	return handler.Up(context.Background())

}

func MigrateDown(ctx *cli.Context) error {
	fmt.Println("Migrating down ...")
	url := fmt.Sprintf("mysql://root@tcp(127.0.0.1:45000)/recipedb")

	handler, err := ConnectDB(url)
	if err != nil {
		return err
	}

	defer handler.Close()

	return handler.Down(context.Background())
}

func ConnectDB(url string) (*migrate.Handle, error) {

	driver, err := mysqlMigrationDriver.Open(url)
	if err != nil {
		return nil, errors.New("Cannot connect to DB, " + err.Error())
	}

	start := time.Now()
	path := "/migrations"
	handler, err := migrate.New(driver, path, migrate.WithHooks(
		func(f file.File) error {
			fmt.Printf("%5d %-60s", f.Version, f.Name)
			start = time.Now()
			return nil
		},
		func(f file.File) error {
			fmt.Printf("completed in %.2fs\n", time.Since(start).Seconds())
			return nil
		},
	))

	if err != nil {
		fmt.Printf("failed in %.2fs\n", time.Since(start).Seconds())
		return nil, err
	}

	return handler, nil

}
