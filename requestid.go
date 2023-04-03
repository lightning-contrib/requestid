package requestid

import "github.com/go-labx/lightning"
import "github.com/matoous/go-nanoid/v2"

type config struct {
	headerKey string
	alphabet  string
	size      int
}

type Options func(*config)

func WithHeaderKey(key string) Options {
	return func(c *config) {
		c.headerKey = key
	}
}

func WithAlphabet(alphabet string) Options {
	return func(c *config) {
		c.alphabet = alphabet
	}
}

func WithSize(size int) Options {
	return func(c *config) {
		c.size = size
	}
}

func Default() lightning.Middleware {
	return New()
}

func New(options ...Options) lightning.Middleware {
	cfg := &config{
		headerKey: "X-Request-ID",
		alphabet:  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		size:      24,
	}

	for _, option := range options {
		option(cfg)
	}

	return func(ctx *lightning.Context) {
		requestId := ctx.Header(cfg.headerKey)
		if requestId == "" {
			id, err := gonanoid.Generate(cfg.alphabet, cfg.size)
			if err == nil {
				requestId = id
			}
		}
		ctx.SetHeader(cfg.headerKey, requestId)
		ctx.Next()
	}
}
