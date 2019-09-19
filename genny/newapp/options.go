package newapp

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/meta"
	"github.com/gobuffalo/pop"
)

type Options struct {
	App     meta.App
	Dialect string
	Prefix  string
}

func (opts *Options) Validate() error {
	if opts.Prefix == "" {
		return fmt.Errorf("you must provide a database name prefix")
	}
	if opts.App.IsZero() {
		opts.App = meta.New(".")
	}

	if len(opts.Dialect) == 0 {
		return fmt.Errorf("you must provide a dialect [%s]", strings.Join(pop.AvailableDialects, ", "))
	}

	var found bool
	for _, d := range pop.AvailableDialects {
		if d == opts.Dialect {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("unknown dialect %q expecting one of %s", opts.Dialect, strings.Join(pop.AvailableDialects, ", "))
	}
	return nil
}
