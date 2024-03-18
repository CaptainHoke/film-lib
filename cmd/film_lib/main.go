package main

import (
	"context"
	actorservice "film-lib/gen/actor_service"
	filmservice "film-lib/gen/film_service"
	searchservice "film-lib/gen/search_service"
	signin "film-lib/gen/sign_in"
	services2 "film-lib/internal/services"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[filmlib] ", log.Ltime)
	}

	// Initialize the services.
	var (
		actorServiceSvc  actorservice.Service
		filmServiceSvc   filmservice.Service
		searchServiceSvc searchservice.Service
		signInSvc        signin.Service
	)
	{
		actorServiceSvc = services2.NewActorService(logger)
		filmServiceSvc = services2.NewFilmService(logger)
		searchServiceSvc = services2.NewSearchService(logger)
		signInSvc = services2.NewSignIn(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		actorServiceEndpoints  *actorservice.Endpoints
		filmServiceEndpoints   *filmservice.Endpoints
		searchServiceEndpoints *searchservice.Endpoints
		signInEndpoints        *signin.Endpoints
	)
	{
		actorServiceEndpoints = actorservice.NewEndpoints(actorServiceSvc)
		filmServiceEndpoints = filmservice.NewEndpoints(filmServiceSvc)
		searchServiceEndpoints = searchservice.NewEndpoints(searchServiceSvc)
		signInEndpoints = signin.NewEndpoints(signInSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:3239/api/v1"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, actorServiceEndpoints, filmServiceEndpoints, searchServiceEndpoints, signInEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
