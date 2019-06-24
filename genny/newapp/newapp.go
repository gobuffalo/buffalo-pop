package newapp

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/gogen"
	"github.com/gobuffalo/genny/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/genny/config"
)

var AvailableDialects = pop.AvailableDialects

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, err
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
		return gg, err
	}
	gg.Add(g)
	return gg, nil
}
