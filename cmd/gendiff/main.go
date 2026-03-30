package main

import (
	"code"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() != 2 {
				return errors.New("number of params != 2")
			}

			filepath1 := cmd.Args().First()
			filepath2 := cmd.Args().Get(1)

			format := cmd.String("format")

			return genDiff(filepath1, filepath2, format)
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func genDiff(filepath1, filepath2, format string) error {
	diff, err := code.GenDiff(filepath1, filepath2, format)
	if err != nil {
		return err
	}

	fmt.Println(diff)

	return nil
}
