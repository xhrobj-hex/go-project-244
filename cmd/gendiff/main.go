package main

import (
	"code"
	"context"
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
				return fmt.Errorf("expected 2 arguments, got %d", cmd.Args().Len())
			}

			filePath1 := cmd.Args().First()
			filePath2 := cmd.Args().Get(1)

			format := cmd.String("format")

			return genDiff(filePath1, filePath2, format)
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func genDiff(leftPath, rightPath, format string) error {
	diff, err := code.GenDiff(leftPath, rightPath, format)
	if err != nil {
		return err
	}

	fmt.Println(diff)

	return nil
}
