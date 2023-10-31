// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cors provides cross-origin resource sharing (CORS) constants for
// Connect. These constants are used to configure the CORS headers for a
// Connect server.
package cors

import "net/http"

// AllowedMethods returns the allowed HTTP methods that scripts running in the
// browser are permitted to use.
//
// To support cross-domain requests with the protocols supported by Connect,
// these headers fields must be included in the preflight response header
// Access-Control-Allow-Methods.
func AllowedMethods() []string {
	return []string{
		http.MethodGet,  // Required for Connect GET requests
		http.MethodPost, // Required for all protocols
	}
}

// AllowedHeaders returns the allowed header fields that scripts running in the
// browser are permitted to access.
//
// To support cross-domain requests with the protocols supported by Connect,
// these field names must be included in header Access-Control-Allow-Headers
// of the actual response.
func AllowedHeaders() []string {
	return []string{
		"Content-Type",             // Required for Connect
		"Connect-Protocol-Version", // Required for Connect
		"Connect-Timeout-Ms",       // Optional for Connect
		"Connect-Accept-Encoding",  // Future use for Connect
		"Connect-Content-Encoding", // Future use for Connect
		"Accept-Encoding",          // Future use for Connect
		"Content-Encoding",         // Future use for Connect
		"Grpc-Timeout",             // Required for gRPC-web
		"X-Grpc-Web",               // Optional for gRPC-web
		"X-User-Agent",             // Optional for gRPC-web
	}
}

// ExposedHeaders returns the headers that scripts running in the browser are
// permitted to see.
//
// To support cross-domain requests with the protocols supported by Connect,
// these field names must be included in header Access-Control-Expose-Headers
// of the actual response.
//
// Make sure to include any application-specific headers your browser client
// should see. If your application uses trailers, they will be sent as header
// fields with a `Trailer-` prefix for Connect unary RPCs - make sure to
// expose them as well if you want them to be visible in all supported
// protocols.
func ExposedHeaders() []string {
	return []string{
		"Content-Encoding",         // Future use for Connect
		"Connect-Content-Encoding", // Future use for Connect
		"Grpc-Status",              // Required for gRPC-web header response
		"Grpc-Message",             // Required for gRPC-web header response
		"Grpc-Status-Details-Bin",  // Required for gRPC-web error details
	}
}
