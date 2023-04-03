package requestid

import "github.com/go-labx/lightning"
import "github.com/matoous/go-nanoid/v2"

func RequestID() lightning.Middleware {
	return func(ctx *lightning.Context) {
		requestId := ctx.Header("x-request-id")
		if requestId == "" {
			id, err := gonanoid.New(24)
			if err == nil {
				requestId = id
			}
		}
		ctx.SetHeader("X-Request-ID", requestId)
		ctx.Next()
	}
}
