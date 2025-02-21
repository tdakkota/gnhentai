// Code generated by ogen, DO NOT EDIT.

package nhentaiapi

import (
	"net/http"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

type (
	optionFunc[C any] func(*C)
)

// ErrorHandler is error handler.
type ErrorHandler = ogenerrors.ErrorHandler

type serverConfig struct {
	NotFound           http.HandlerFunc
	MethodNotAllowed   func(w http.ResponseWriter, r *http.Request, allowed string)
	ErrorHandler       ErrorHandler
	Prefix             string
	Middleware         Middleware
	MaxMultipartMemory int64
}

// ServerOption is server config option.
type ServerOption interface {
	applyServer(*serverConfig)
}

var _ ServerOption = (optionFunc[serverConfig])(nil)

func (o optionFunc[C]) applyServer(c *C) {
	o(c)
}

func newServerConfig(opts ...ServerOption) serverConfig {
	cfg := serverConfig{
		NotFound: http.NotFound,
		MethodNotAllowed: func(w http.ResponseWriter, r *http.Request, allowed string) {
			status := http.StatusMethodNotAllowed
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Methods", allowed)
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				status = http.StatusNoContent
			} else {
				w.Header().Set("Allow", allowed)
			}
			w.WriteHeader(status)
		},
		ErrorHandler:       ogenerrors.DefaultErrorHandler,
		Middleware:         nil,
		MaxMultipartMemory: 32 << 20, // 32 MB
	}
	for _, opt := range opts {
		opt.applyServer(&cfg)
	}
	return cfg
}

type baseServer struct {
	cfg serverConfig
}

func (s baseServer) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

func (s baseServer) notAllowed(w http.ResponseWriter, r *http.Request, allowed string) {
	s.cfg.MethodNotAllowed(w, r, allowed)
}

func (cfg serverConfig) baseServer() (s baseServer, err error) {
	s = baseServer{cfg: cfg}
	return s, nil
}

type clientConfig struct {
	Client ht.Client
}

// ClientOption is client config option.
type ClientOption interface {
	applyClient(*clientConfig)
}

var _ ClientOption = (optionFunc[clientConfig])(nil)

func (o optionFunc[C]) applyClient(c *C) {
	o(c)
}

func newClientConfig(opts ...ClientOption) clientConfig {
	cfg := clientConfig{
		Client: http.DefaultClient,
	}
	for _, opt := range opts {
		opt.applyClient(&cfg)
	}
	return cfg
}

type baseClient struct {
	cfg clientConfig
}

func (cfg clientConfig) baseClient() (c baseClient, err error) {
	c = baseClient{cfg: cfg}
	return c, nil
}

// Option is config option.
type Option interface {
	ServerOption
	ClientOption
}

// WithClient specifies http client to use.
func WithClient(client ht.Client) ClientOption {
	return optionFunc[clientConfig](func(cfg *clientConfig) {
		if client != nil {
			cfg.Client = client
		}
	})
}

// WithNotFound specifies Not Found handler to use.
func WithNotFound(notFound http.HandlerFunc) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if notFound != nil {
			cfg.NotFound = notFound
		}
	})
}

// WithMethodNotAllowed specifies Method Not Allowed handler to use.
func WithMethodNotAllowed(methodNotAllowed func(w http.ResponseWriter, r *http.Request, allowed string)) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if methodNotAllowed != nil {
			cfg.MethodNotAllowed = methodNotAllowed
		}
	})
}

// WithErrorHandler specifies error handler to use.
func WithErrorHandler(h ErrorHandler) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if h != nil {
			cfg.ErrorHandler = h
		}
	})
}

// WithPathPrefix specifies server path prefix.
func WithPathPrefix(prefix string) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		cfg.Prefix = prefix
	})
}

// WithMiddleware specifies middlewares to use.
func WithMiddleware(m ...Middleware) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		switch len(m) {
		case 0:
			cfg.Middleware = nil
		case 1:
			cfg.Middleware = m[0]
		default:
			cfg.Middleware = middleware.ChainMiddlewares(m...)
		}
	})
}

// WithMaxMultipartMemory specifies limit of memory for storing file parts.
// File parts which can't be stored in memory will be stored on disk in temporary files.
func WithMaxMultipartMemory(max int64) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if max > 0 {
			cfg.MaxMultipartMemory = max
		}
	})
}
