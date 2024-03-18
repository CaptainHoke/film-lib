package main

import (
	"context"
	actorservice "film-lib/gen/actor_service"
	filmservice "film-lib/gen/film_service"
	actorservicesvr "film-lib/gen/http/actor_service/server"
	filmservicesvr "film-lib/gen/http/film_service/server"
	searchservicesvr "film-lib/gen/http/search_service/server"
	signinsvr "film-lib/gen/http/sign_in/server"
	searchservice "film-lib/gen/search_service"
	signin "film-lib/gen/sign_in"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, actorServiceEndpoints *actorservice.Endpoints, filmServiceEndpoints *filmservice.Endpoints, searchServiceEndpoints *searchservice.Endpoints, signInEndpoints *signin.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		actorServiceServer  *actorservicesvr.Server
		filmServiceServer   *filmservicesvr.Server
		searchServiceServer *searchservicesvr.Server
		signInServer        *signinsvr.Server
	)
	{
		eh := errorHandler(logger)
		actorServiceServer = actorservicesvr.New(actorServiceEndpoints, mux, dec, enc, eh, nil)
		filmServiceServer = filmservicesvr.New(filmServiceEndpoints, mux, dec, enc, eh, nil)
		searchServiceServer = searchservicesvr.New(searchServiceEndpoints, mux, dec, enc, eh, nil)
		signInServer = signinsvr.New(signInEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				actorServiceServer,
				filmServiceServer,
				searchServiceServer,
				signInServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	actorservicesvr.Mount(mux, actorServiceServer)
	filmservicesvr.Mount(mux, filmServiceServer)
	searchservicesvr.Mount(mux, searchServiceServer)
	signinsvr.Mount(mux, signInServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range actorServiceServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range filmServiceServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range searchServiceServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range signInServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Printf("failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
