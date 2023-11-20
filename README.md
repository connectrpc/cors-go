cors-go
===============

Cross-origin resource sharing (CORS) support for Connect servers. Exports methods
to configure CORS headers to allow Connect and gRPC-web protocols to operate in
the browser.

## Example

As an example, we will use the [github.com/rs/cors](https://github.com/rs/cors) 
package to demonstrate how to use the constants defined in this package.

```go
import (
	cors "github.com/bufbuild/cors-go"
	rscors "github.com/rs/cors"
)

// corsMiddleware wraps a handler with require cors config using rs/cors.
func corsMiddleware(handler http.Handler) http.Handler {
	// Create a new cors instance with default options.
	c := rscors.New(rscors.Options{
		AllowedMethods: cors.AllowedMethods(),
		AllowedHeaders: cors.AllowedHeaders(),
		ExposedHeaders: cors.ExposedHeaders(),
	})
	// Insert the middleware as a wrapper around your handler.
	return c.Handler(handler /* connect handler */)
}
```

## Status: Alpha

Cors is undergoing initial development and is not yet stable.

## Legal

Offered under the [Apache 2 license][license].

[license]: https://github.com/bufbuild/cors-go/blob/main/LICENSE
