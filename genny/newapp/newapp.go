package newapp

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/gogen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/genny/config"
	"github.com/pkg/errors"
)

var AvailableDialects = pop.AvailableDialects

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, errors.WithStack(err)
	}

	g := genny.New()
	g.Box(packr.New("github.com/gobuffalo/buffalo-pop/genny/newapp/templates", "../newapp/templates"))

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())

	g.Command(gogen.Get("github.com/gobuffalo/pop"))

	gg.Add(g)

	g, err := config.New(&config.Options{
		Dialect:  opts.Dialect,
		Prefix:   opts.Prefix,
		FileName: "database.yml",
		Root:     opts.App.Root,
	})
	if err != nil {
		return gg, errors.WithStack(err)
	}
	gg.Add(g)
	return gg, nil
}
