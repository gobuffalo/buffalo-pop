package newapp

import (
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush/v4"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/pop/v5/genny/config"
)

var AvailableDialects = pop.AvailableDialects

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, err
	}

	g := genny.New()
	g.Box(packr.New("github.com/gobuffalo/buffalo-pop/v2/genny/newapp/templates", "../newapp/templates"))

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())

	gg.Add(g)

	g, err := config.New(&config.Options{
		Dialect:  opts.Dialect,
		Prefix:   opts.Prefix,
		FileName: "database.yml",
		Root:     opts.Root,
	})
	if err != nil {
		return gg, err
	}
	gg.Add(g)
	return gg, nil
}
