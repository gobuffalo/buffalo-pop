package newapp

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/pop/v5"
)

type Options struct {
	Dialect string
	Prefix  string
}

func (opts *Options) Validate() error {
	if opts.Prefix == "" {
		return fmt.Errorf("you must provide a database name prefix")
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
