# requestid

The `Lightning` Request ID middleware is designed to provide unique request IDs for every incoming request and add them to logs and response headers. This helps with tracking request in a distributed system, making it easier to debug and troubleshoot issues.

## Installation

To use the Request ID middleware, you can install it using the following command:

```bash
go get github.com/lightning-contrib/requestid
```

## Usage

Here is an example of how to use `requestid` middleware:

```go
package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/requestid"
)

func main() {
	app := lightning.NewApp()
	app.Use(requestid.Default())

	app.Get("/ping", func(ctx *lightning.Context) {
		ctx.JSON(200, map[string]string{
			"message": "pong",
		})
	})

	app.Run()
}
```

You can also use the `New()` function, it also returns a middleware, but it allows for customization of the config struct through the use of variadic Options arguments.

```go
app.Use(requestid.New(
    requestid.WithAlphabet("1234567890"),
    requestid.WithHeaderKey("X-Custom-ID"),
    requestid.WithSize(16)),
)
```

## API Documentation

For detailed API documentation and usage examples, please refer to the [documentation](https://pkg.go.dev/github.com/lightning-contrib/requestid).

## Contributing

If you'd like to contribute to lightning, please
see [CONTRIBUTING.md](https://github.com/lightning-contrib/requestid/blob/main/CONTRIBUTING.md) for guidelines.

## License

lightning is licensed under the [MIT License](https://github.com/lightning-contrib/requestid/blob/main/LICENSE).
