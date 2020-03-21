package server

import (
	"github.com/rs/zerolog"
)

type Option func(c *Server)

func WithLogger(logger zerolog.Logger) Option {
	return func(c *Server) {
		c.log = logger
	}
}

func WithCache(cache Cache) Option {
	return func(c *Server) {
		c.cache = cache
	}
}
