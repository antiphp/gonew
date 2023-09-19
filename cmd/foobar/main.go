// Package main runs the agent.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ettle/strcase"
	"github.com/hamba/cmd/v2"
	"github.com/urfave/cli/v2"
)

const (
	flagFoobar = "foobar"

	catFoobar = "Foobar"
)

var (
	buildVersion   = "<unknown>"
	buildTimestamp = "0"
	buildTime      = time.Unix(func() int64 { n, _ := strconv.Atoi(buildTimestamp); return int64(n) }(), 0)
)

var flags = cmd.Flags{
	&cli.IntFlag{
		Name:     flagFoobar,
		Usage:    "Sets foobar.",
		Category: catFoobar,
		EnvVars:  []string{strcase.ToSNAKE(flagFoobar)},
	},
}.Merge(cmd.LogFlags)

func main() {
	app := cli.NewApp()
	app.Name = "Foobar"
	app.Version = buildVersion + " @ " + buildTime.Format(time.RFC3339)
	app.Flags = flags
	app.Action = run

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := app.RunContext(ctx, os.Args); err != nil {
		switch { //nolint:gocritic // Be flexible.
		case !errors.Is(err, errAlreadyLogged):
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
		os.Exit(1)
	}
}
