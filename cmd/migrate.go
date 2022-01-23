package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Migrate(ctx *cli.Context) error {
	fmt.Println("Migrating....")
	fmt.Println(ctx.Args())
	return nil
}
