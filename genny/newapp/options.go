package newapp

import (
	"github.com/gobuffalo/meta"
	"github.com/pkg/errors"
)

type Options struct {
	App     meta.App
	Dialect string
	Prefix  string
}

func (opts *Options) Validate() error {
	if opts.Prefix == "" {
		return errors.New("you must provide a database name prefix")
	}
	if opts.App.IsZero() {
		opts.App = meta.New(".")
	}

	return nil
}
