package main

import (
	"context"
	actorservice "film-lib/gen/actor_service"
	filmservice "film-lib/gen/film_service"
	isalive "film-lib/gen/is_alive"
	searchservice "film-lib/gen/search_service"
	signin "film-lib/gen/sign_in"
	"film-lib/internal/repo"
	services2 "film-lib/internal/services"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
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

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=10",
		os.Getenv("FL_DBUSER"),
		os.Getenv("FL_DBPASS"),
		os.Getenv("FL_DBHOST"),
		os.Getenv("FL_DBPORT"),
		os.Getenv("FL_DBNAME"))
	pg, _ := repo.NewPostgresDB(context.Background(), connString)

	if err := pg.Ping(context.Background()); err != nil {
		logger.Fatal("Failed to ping db")
	}

	actorsTableExists, _ := pg.TableExists(context.Background(), "actors")
	filmsTableExists, _ := pg.TableExists(context.Background(), "films")
	actorFilmTableExists, _ := pg.TableExists(context.Background(), "actor_film")
	shouldCreateTables := !actorsTableExists || !filmsTableExists || !actorFilmTableExists

	// TODO: migrate package or smth, this is broken
	if shouldCreateTables {
		migrationsDir := "./db/migrations"
		items, _ := os.ReadDir(migrationsDir)

		for _, item := range items {
			if !item.IsDir() {
				path := filepath.Join(migrationsDir, item.Name())
				m, _ := os.ReadFile(path)
				sql := string(m)
				logger.Println(sql)
				_, err := pg.Db.Exec(context.Background(), sql)
				if err != nil {
					logger.Fatal("Failed to set up db")
				}
			}
		}
	}

	// Initialize the services.
	var (
		actorServiceSvc  actorservice.Service
		filmServiceSvc   filmservice.Service
		searchServiceSvc searchservice.Service
		signInSvc        signin.Service
		isAliveSvc       isalive.Service
	)
	{
		actorServiceSvc = services2.NewActorService(logger)
		filmServiceSvc = services2.NewFilmService(logger)
		searchServiceSvc = services2.NewSearchService(logger)
		signInSvc = services2.NewSignIn(logger)
		isAliveSvc = services2.NewIsAliveService(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		actorServiceEndpoints  *actorservice.Endpoints
		filmServiceEndpoints   *filmservice.Endpoints
		searchServiceEndpoints *searchservice.Endpoints
		signInEndpoints        *signin.Endpoints
		isAliveEndPoints       *isalive.Endpoints
	)
	{
		actorServiceEndpoints = actorservice.NewEndpoints(actorServiceSvc)
		filmServiceEndpoints = filmservice.NewEndpoints(filmServiceSvc)
		searchServiceEndpoints = searchservice.NewEndpoints(searchServiceSvc)
		signInEndpoints = signin.NewEndpoints(signInSvc)
		isAliveEndPoints = isalive.NewEndpoints(isAliveSvc)
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
	addr := ""
	switch *hostF {
	// TODO: Should've been fixed in api.go
	case "docker":
		{
			addr = "http://0.0.0.0:3239/api/v1"
		}
	case "localhost":
		{
			addr = "http://localhost:3239/api/v1"
		}
	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: docker, localhost)\n", *hostF)
	}

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
	handleHTTPServer(ctx, u, actorServiceEndpoints, filmServiceEndpoints, searchServiceEndpoints, signInEndpoints, isAliveEndPoints, &wg, errc, logger, *dbgF)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
