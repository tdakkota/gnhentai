package main

import (
	"encoding/json"
	"fmt"
	"github.com/tdakkota/gnhentai"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type App struct {
	client *gnhentai.Client
}

func NewApp() *App {
	return &App{
		client: gnhentai.NewClient(),
	}
}

func prettyPrint(d gnhentai.Doujinshi) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func (app *App) random(c *cli.Context) error {
	doujinshi, err := app.client.Random()
	if err != nil {
		return err
	}

	err = prettyPrint(doujinshi)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) get(c *cli.Context) error {
	doujinshi, err := app.client.ByID(c.Int("id"))
	if err != nil {
		return err
	}

	err = prettyPrint(doujinshi)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "random",
			Action: app.random,
		},
		{
			Name: "get",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "id",
					Usage: "id of doujinshi",
				},
			},
			Action: app.get,
		},
	}
}

func (app *App) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "proxy",
			Usage:   "proxy for http client (in format socks5://localhost:9051)",
			EnvVars: []string{"PROXY"},
		},
	}
}

func (app *App) cli() *cli.App {
	return &cli.App{
		Name:     "gnhentai-cli",
		Usage:    "cli tool to search Doujinshi",
		Commands: app.commands(),
		Flags:    app.flags(),
	}

}

func (app *App) Run(args []string) error {
	return app.cli().Run(args)
}

func main() {
	if err := NewApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
