// Copyright 2017 Improbable. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcweb

import "net/http"

var (
	defaultOptions = &options{
		allowedRequestHeaders:          []string{"*"},
		corsForRegisteredEndpointsOnly: true,
		originFunc:                     func(origin string) bool { return true },
	}
)

type options struct {
	allowedRequestHeaders          []string
	corsForRegisteredEndpointsOnly bool
	originFunc                     func(origin string) bool
	enableWebsockets               bool
	websocketOriginFunc            func(req *http.Request) bool
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

// WithOriginFunc allows for customizing what CORS Origin requests are allowed.
//
// This is controlling the CORS pre-flight `Access-Control-Allow-Origin`. This mechanism allows you to limit the
// availability of the APIs based on the domain name of the calling website (Origin). You can provide a function that
// filters the allowed Origin values.
//
// The default behaviour is `*`, i.e. to allow all calling websites.
//
// The relevant CORS pre-flight docs:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
func WithOriginFunc(originFunc func(origin string) bool) Option {
	return func(o *options) {
		o.originFunc = originFunc
	}
}

// WithCorsForRegisteredEndpointsOnly allows for customizing whether OPTIONS requests with the `X-GRPC-WEB` header will
// only be accepted if they match a registered gRPC endpoint.
//
// This should be set to false to allow handling gRPC requests for unknown endpoints (e.g. for proxying).
//
// The default behaviour is `true`, i.e. only allows CORS requests for registered endpoints.
func WithCorsForRegisteredEndpointsOnly(onlyRegistered bool) Option {
	return func(o *options) {
		o.corsForRegisteredEndpointsOnly = onlyRegistered
	}
}

// WithAllowedRequestHeaders allows for customizing what gRPC request headers a browser can add.
//
// This is controlling the CORS pre-flight `Access-Control-Allow-Headers` method and applies to *all* gRPC handlers.
// However, a special `*` value can be passed in that allows
// the browser client to provide *any* header, by explicitly whitelisting all `Access-Control-Request-Headers` of the
// pre-flight request.
//
// The default behaviour is `[]string{'*'}`, allowing all browser client headers. This option overrides that default,
// while maintaining a whitelist for gRPC-internal headers.
//
// Unfortunately, since the CORS pre-flight happens independently from gRPC handler execution, it is impossible to
// automatically discover it from the gRPC handler itself.
//
// The relevant CORS pre-flight docs:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
func WithAllowedRequestHeaders(headers []string) Option {
	return func(o *options) {
		o.allowedRequestHeaders = headers
	}
}

// WithWebsockets allows for handling grpc-web requests of websockets - enabling bidirectional requests.
//
// The default behaviour is false, i.e. to disallow websockets
func WithWebsockets(enableWebsockets bool) Option {
	return func(o *options) {
		o.enableWebsockets = enableWebsockets
	}
}

// WithWebsocketOriginFunc allows for customizing the acceptance of Websocket requests - usually to check that the origin
// is valid.
//
// The default behaviour is to check that the origin of the request matches the host of the request.
func WithWebsocketOriginFunc(websocketOriginFunc func(req *http.Request) bool) Option {
	return func(o *options) {
		o.websocketOriginFunc = websocketOriginFunc
	}
}