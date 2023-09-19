// Package cmd contains reusable, command-related helper functions.
package cmd

import (
	"io"

	"github.com/hamba/cmd/v2"
	"github.com/hamba/logger/v2"
	lctx "github.com/hamba/logger/v2/ctx"
	"github.com/urfave/cli/v2"
)

// NewLogger returns a preconfigured logger.
func NewLogger(c *cli.Context, w io.Writer, svc string) (*logger.Logger, func(), error) {
	logger.TimeFormat = logger.TimeFormatISO8601

	log, err := newLogger(c, w)
	if err != nil {
		return nil, nil, err
	}
	cancel := log.WithTimestamp()

	log = log.With(lctx.Str("svc", svc))

	return log, cancel, nil
}

func newLogger(c *cli.Context, w io.Writer) (*logger.Logger, error) {
	str := c.String(cmd.FlagLogLevel)
	if str == "" {
		str = "info"
	}

	lvl, err := logger.LevelFromString(str)
	if err != nil {
		return nil, err
	}

	fmtr := newLogFormatter(c)

	tags, err := cmd.Split(c.StringSlice(cmd.FlagLogCtx), "=")
	if err != nil {
		return nil, err
	}

	fields := make([]logger.Field, len(tags))
	for i, t := range tags {
		fields[i] = lctx.Str(t[0], t[1])
	}

	return logger.New(w, fmtr, lvl).With(fields...), nil
}

func newLogFormatter(c *cli.Context) logger.Formatter {
	format := c.String(cmd.FlagLogFormat)
	switch format {
	case "json":
		return logger.JSONFormat()
	case "console":
		return logger.ConsoleFormat()
	default:
		return logger.LogfmtFormat()
	}
}
