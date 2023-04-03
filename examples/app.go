package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/requestid"
)

func main() {
	app := lightning.NewApp()

	app.Use(requestid.New(
		requestid.WithAlphabet("1234567890"),
		requestid.WithHeaderKey("X-Custom-ID"),
		requestid.WithSize(16)),
	)

	app.Get("/ping", func(ctx *lightning.Context) {
		ctx.JSON(200, map[string]string{
			"message": "pong",
		})
	})

	app.Run()
}
