package newapp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Options(t *testing.T) {
	r := require.New(t)

	opts := &Options{}
	r.Error(opts.Validate())

	opts.Dialect = "foo"
	r.Error(opts.Validate())

	opts.Dialect = "postgres"
	r.Error(opts.Validate())

	opts.Prefix = "coke"
	r.NoError(opts.Validate())

	opts.Dialect = "asdf"
	r.Error(opts.Validate())

	opts.Dialect = "postgres"
	r.NoError(opts.Validate())
}
