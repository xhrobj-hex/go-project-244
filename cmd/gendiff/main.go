package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"code/internal/parser"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},

		// https://cli.urfave.org/v3/examples/arguments/basics/

		Action: func(ctx context.Context, cmd *cli.Command) error {
			for i := 0; i < cmd.Args().Len(); i++ {
				filepath := cmd.Args().Get(i)
				json, err := parser.Parse(filepath)
				if err != nil {
					return err
				}
				fmt.Printf("%#v\n", json)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
