package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func runner(c *cli.Context) error {
	return nil
}

func commands() []*cli.Command {
	return []*cli.Command{
		{Name: "random"},
		{Name: "get"},
	}
}

func main() {
	app := &cli.App{
		Name:   "gnhentai-cli",
		Usage:  "cli tool to search Doujinshi",
		Action: runner,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
