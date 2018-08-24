package newapp

import (
	"context"
	"os"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/pop"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	for _, d := range pop.AvailableDialects {
		run := genny.DryRunner(context.Background())
		err := run.Chdir(os.TempDir(), func() error {
			g, err := New(&Options{
				Prefix:  "foo",
				Dialect: d,
			})
			r.NoError(err)
			run.With(g)

			r.NoError(run.Run())

			res := run.Results()
			r.Len(res.Commands, 1)
			r.Len(res.Files, 3)
			return nil
		})
		r.NoError(err)
	}
}
