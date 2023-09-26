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
	flagAddr = "addr"

	catFoobar = "Foobar"
)

var (
	buildVersion   = "<unknown>"
	buildTimestamp = "0"
	buildTime      = time.Unix(func() int64 { n, _ := strconv.Atoi(buildTimestamp); return int64(n) }(), 0)
)

var flags = cmd.Flags{
	&cli.StringFlag{
		Name:     flagAddr,
		Usage:    "HTTP address to listen to",
		Category: catFoobar,
		Value:    ":8080",
		EnvVars:  []string{strcase.ToSNAKE(flagAddr)},
	},
}.Merge(cmd.LogFlags, cmd.StatsFlags, cmd.TracingFlags)

func main() {
	os.Exit(mainWithExitCode())
}

func mainWithExitCode() int {
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
		return 1
	}
	return 0
}
