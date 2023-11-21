cors
====

[![Build](https://github.com/connectrpc.com/cors-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/connectrpc.com/cors-go/actions/workflows/ci.yaml)
[![Report Card](https://goreportcard.com/badge/connectrpc.com/cors)](https://goreportcard.com/report/connectrpc.com/cors)
[![GoDoc](https://pkg.go.dev/badge/connectrpc.com/cors.svg)](https://pkg.go.dev/connectrpc.com/cors)

`connectrpc.com/cors` provides convenience methods to make configuring
Cross-Origin Resource Sharing (CORS) easier for
[Connect](https://github.com/connectrpc/connect-go) servers. CORS is often
required for the Connect and gRPC-Web protocols to work correctly in web
browsers.

For background, more details, and best practices, see [Connect's CORS
documentation](https://connectrpc.com/docs/cors).

## Example

This package should work with any CORS package. As an example, we'll use it
with [github.com/rs/cors](https://github.com/rs/cors).

```go
import (
	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://acme.com"}, // replace with your domain
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
    MaxAge: 7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}
```

## Status: Alpha

This module is undergoing initial development and is not yet stable.

## Legal

Offered under the [Apache 2 license][LICENSE].
