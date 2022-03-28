package cmd

import (
	db "github.com/jonDufty/recipes/common/database"
	cli "github.com/urfave/cli/v2"
)

func MigrateUp(ctx *cli.Context) error {
	return db.MigrateUp(ctx)
}

func MigrateDown(ctx *cli.Context) error {
	return db.MigrateDown(ctx)
}
