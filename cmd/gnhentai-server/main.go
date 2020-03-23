package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/api"
	"github.com/tdakkota/gnhentai/server"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

type App struct {
	client     gnhentai.Client
	downloader gnhentai.Downloader
}

func NewApp() *App {
	return &App{}
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

	apiClient := api.NewClient(api.WithClient(client))
	app.client = apiClient
	app.downloader = apiClient
	return nil
}

func (app *App) run(c *cli.Context) error {
	r := app.setupServer(c)
	bind := c.String("server.bind")
	log.Info().Str("addr", bind).Msgf("API server listening on %s", bind)
	return http.ListenAndServe(bind, r)
}

func (app *App) setupServer(c *cli.Context) http.Handler {
	r := chi.NewRouter()

	logger := zerolog.New(os.Stdout)
	r.Use(
		server.Logger(logger),
		middleware.Timeout(c.Duration("server.timeout")),
	)

	server.NewServer(
		app.client,
		app.downloader,
		server.WithLogger(logger),
	).Register(r)

	return r
}

func (app *App) commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "run",
			Description: "runs server",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "server.bind",
					Required: false,
					Value:    ":8080",
					Usage:    "addr to bind",
				},
				&cli.DurationFlag{
					Name:     "server.timeout",
					Required: false,
					Value:    30 * time.Second,
					Usage:    "server connection timeout",
				},
			},
			Action: app.run,
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
	}
}

func (app *App) cli() *cli.App {
	return &cli.App{
		Name:     "gnhentai-server",
		Usage:    "simple API server like nhentai.net",
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
		log.Error().Err(err).Msg("failed")
		os.Exit(-1)
	}
}
