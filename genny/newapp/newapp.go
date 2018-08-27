package newapp

import (
	"errors"
	"os/exec"

	"github.com/gobuffalo/buffalo/meta"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/pop/soda/cmd/generate"
)

func New(opts *Options) (*genny.Generator, error) {
	if opts.Prefix == "" {
		return nil, errors.New("you must provide a database name prefix")
	}
	if (opts.App == meta.App{}) {
		opts.App = meta.New(".")
	}
	g := genny.New()
	g.Box(packr.NewBox("../newapp/templates"))

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Replace("-dot-", "."))
	g.Command(exec.Command(genny.GoBin(), "get", "github.com/gobuffalo/pop/..."))
	g.RunFn(func(r *genny.Runner) error {
		data := map[string]interface{}{
			"dialect": opts.Dialect,
			"name":    opts.Prefix,
		}
		return generate.Config("./database.yml", data)
	})
	return g, nil
}
