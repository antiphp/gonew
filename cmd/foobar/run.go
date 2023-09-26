package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/antiphp/gonew"
	"github.com/antiphp/gonew/api"
	"github.com/antiphp/gonew/cmd"
	httpx "github.com/antiphp/gonew/internal/http"
	lctx "github.com/hamba/logger/v2/ctx"
	"github.com/urfave/cli/v2"
)

const svc = "foobar"

// errAlreadyLogged represents an error in the run function and allows main to determine whether the error has already been
// logged (and its content can be ignored) or not. This is due to the fact that the logger is created inside run.
var errAlreadyLogged = errors.New("run error")

func run(c *cli.Context) error {
	ctx, cancel := context.WithCancel(c.Context)
	defer cancel()

	log, logClose, err := cmd.NewLogger(c, svc)
	if err != nil {
		return fmt.Errorf("creating logger: %w", err)
	}
	defer logClose()

	app, err := gonew.New(log)
	if err != nil {
		log.Error("Could not create app", lctx.Err(err))
		return errAlreadyLogged
	}

	apiSrv := api.NewServer(app, log)

	mux := http.NewServeMux()
	mux.Handle("/ready", httpx.OKHandler())
	mux.Handle("/live", httpx.OKHandler())
	mux.Handle("/", apiSrv)

	addr := c.String(flagAddr)
	httpSrv := httpx.NewServer(ctx, addr, mux)

	log.Info("Starting foobar", lctx.Str("buildVersion", c.App.Version), lctx.Str("addr", addr))
	httpSrv.Serve(func(err error) {
		log.Error("Server error", lctx.Err(err))
		cancel()
	})
	defer func() { _ = httpSrv.Close() }()

	<-ctx.Done()

	log.Info("Shutting down")
	if err = httpSrv.Shutdown(10 * time.Second); err != nil {
		log.Error("Failed to shutdown server", lctx.Err(err))
		// Fall through.
	}

	return nil
}
