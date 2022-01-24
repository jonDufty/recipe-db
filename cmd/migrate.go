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

func Migrate(ctx *cli.Context) error {
	fmt.Println("Migrating....")
	fmt.Println(ctx.Args())

	url := "mysql://root@tcp(127.0.0.1:45000)/recipedb"
	driver, err := mysqlMigrationDriver.Open(url)

	if err != nil {
		return errors.New("Cannot connect to DB, " + err.Error())
	}
	start := time.Now()
	path := "/Users/jondufty/Projects/99designs/recipe-db/auth/migrations"
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
		return err
	}

	defer handler.Close()

	return handler.Up(context.Background())
}
