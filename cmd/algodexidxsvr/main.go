package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	info "algodexidx/gen/info"

	algodexidx "algodexidx"
	"algodexidx/cmd/algodexidxsvr/backend"
	account "algodexidx/gen/account"
	inspect "algodexidx/gen/inspect"
)

var GitSummary string

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")

		network = flag.String("network", "", "Algorand network to connect to (testnet or mainnet) - or ALGODEX_NETWORK env.")
	)
	flag.Parse()

	if *network == "" {
		*network = os.Getenv("ALGODEX_NETWORK")
	}
	if *network != "testnet" && *network != "mainnet" {
		fmt.Fprintf(os.Stderr, "invalid network %s\n", *network)
		os.Exit(1)
	}

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[algodexidx] ", log.Ltime)
	}

	// Initialize the services.
	var (
		accountSvc account.Service
		inspectSvc inspect.Service
		infoSvc    info.Service
	)
	{
		accountSvc = algodexidx.NewAccount(logger)
		inspectSvc = algodexidx.NewInspect(logger)
		infoSvc = algodexidx.NewInfo(logger, GitSummary)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		accountEndpoints *account.Endpoints
		inspectEndpoints *inspect.Endpoints
		infoEndpoints    *info.Endpoints
	)
	{
		accountEndpoints = account.NewEndpoints(accountSvc)
		inspectEndpoints = inspect.NewEndpoints(inspectSvc)
		infoEndpoints = info.NewEndpoints(infoSvc)
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

	backend.InitBackend(ctx, logger, *network)

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://:80"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
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
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, ":80")
			}
			handleHTTPServer(ctx, u, accountEndpoints, inspectEndpoints, infoEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}
	logger.Printf("Version:%s", GitSummary)
	logger.Printf("Network:%s", *network)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
