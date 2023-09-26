// Package gonew contains a blueprint project skeleton.
package gonew

import (
	"context"
	"errors"

	"github.com/hamba/logger/v2"
)

// Foobar represents a foobar application.
type Foobar struct {
	rng int

	log *logger.Logger
}

// OptFunc represents an optional foobar configuration.
type OptFunc func(*Foobar)

// WithNotSoRandomOpt applies a not so random foobar configuration option.
func WithNotSoRandomOpt(rng int) OptFunc {
	return func(f *Foobar) {
		f.rng = rng
	}
}

// New returns a new foobar.
func New(log *logger.Logger, opts ...OptFunc) (*Foobar, error) {
	f := Foobar{
		log: log,
	}

	for _, opt := range opts {
		opt(&f)
	}

	return &f, nil
}

// Foobar runs foobar.
func (f *Foobar) Foobar(_ context.Context, msg string) (string, error) {
	if msg == "foobar" {
		return "", errors.New("foobar") //nolint:goerr113 // It's just an example.
	}
	return msg, nil
}
