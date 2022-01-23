package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func InitDB(ctx *cli.Context) error {
	fmt.Println("Hello - InitDB")
	return nil
}
