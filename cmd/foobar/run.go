package main

import (
	"errors"
	"os"

	"github.com/antiphp/gonew/cmd"
	"github.com/urfave/cli/v2"
)

const svc = "foobar"

// errAlreadyLogged represents an error in the run function and allows main to determine whether the error has already been
// logged (and its content can be ignored) or not. This is due to the fact that the logger is created inside run.
var errAlreadyLogged = errors.New("run error")

func run(c *cli.Context) error {
	log, closeLog, err := cmd.NewLogger(c, os.Stdout, svc)
	if err != nil {
		return err
	}
	defer closeLog()

	log.Info("Starting foobar server")

	return nil
}
