package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
	"golang.org/x/net/proxy"

	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/api"
	"github.com/tdakkota/gnhentai/downloaders"
	"github.com/tdakkota/gnhentai/parser"
)

type App struct {
	client     gnhentai.Client
	downloader gnhentai.Downloader
}

func NewApp() *App {
	return &App{}
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

func (app *App) getDoujinshi(c *cli.Context) (doujinshi gnhentai.Doujinshi, dir string, err error) {
	id, dir := c.Int("id"), c.String("dir")

	if id != 0 {
		doujinshi, err = app.client.ByID(id)
	} else {
		doujinshi, err = app.client.Random()
	}
	if err != nil {
		return doujinshi, "", err
	}

	if dir == "" {
		dir = strconv.Itoa(id)
	}

	return
}

func (app *App) download(c *cli.Context) (err error) {
	doujinshi, dir, err := app.getDoujinshi(c)
	if err != nil {
		return err
	}

	fmt.Println("Downloading", doujinshi.Name())
	fmt.Println("Tags:")
	for _, tag := range doujinshi.Tags {
		fmt.Println(" - ", tag.Name)
	}

	err = gnhentai.DownloadAll(app.downloader, doujinshi, func(i int, d gnhentai.Doujinshi) string {
		format := gnhentai.FormatFromImage(doujinshi.Images.Pages[i])
		return filepath.Join(dir, fmt.Sprintf("%d.%s", i, format))
	})
	if err != nil {
		return err
	}

	err = prettyPrint(doujinshi)
	if err != nil {
		return err
	}

	return nil
}

func transportWithSocks(rawurl string) (http.RoundTripper, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	var auth *proxy.Auth

	if p, ok := u.User.Password(); ok {
		auth = &proxy.Auth{Password: p}
	}

	if user := u.User.Username(); user != "" {
		if auth == nil {
			auth = &proxy.Auth{}
		}
		auth.User = user
	}

	dialer, err := proxy.SOCKS5("tcp", u.Host+":"+u.Port(), auth, proxy.Direct)
	if err != nil {
		return nil, err
	}

	httpTransport := &http.Transport{}
	httpTransport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		return dialer.Dial(network, addr)
	}
	return httpTransport, nil
}

func (app *App) setup(c *cli.Context) error {
	client := http.DefaultClient

	if proxyURL := c.String("proxy"); proxyURL != "" {
		transport, err := transportWithSocks(proxyURL)
		if err != nil {
			return err
		}
		client.Transport = transport
	}

	if c.Bool("api") {
		app.client = api.NewClient(api.WithClient(client))
	} else {
		app.client = parser.NewParser(parser.WithClient(client))
	}

	app.downloader = downloaders.NewSimpleDownloader(client)

	return nil
}

func (app *App) commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "random",
			Description: "returns metadata of random doujinshi",
			Action:      app.random,
		},
		{
			Name:        "get",
			Description: "returns metadata of doujinshi by given id",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "id",
					Usage: "id of doujinshi",
				},
			},
			Action: app.get,
		},
		{
			Name:        "download",
			Description: "downloads pages to directory",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "id",
					Usage: "id of doujinshi",
				},
				&cli.StringFlag{
					Name:  "dir",
					Usage: "output directory",
				},
			},
			Action: app.download,
		},
	}
}

func (app *App) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "proxy",
			Usage:   "proxy for http client (in format socks5://localhost:9051)",
			EnvVars: []string{"PROXY", "GNHENTAI_PROXY"},
		},
		&cli.BoolFlag{
			Name:    "api",
			Usage:   "if true, uses api.Client instead of parser",
			EnvVars: []string{"API", "GNHENTAI_API"},
		},
	}
}

func (app *App) cli() *cli.App {
	return &cli.App{
		Name:     "gnhentai-cli",
		Usage:    "cli tool to search Doujinshi",
		Before:   app.setup,
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
