package newapp

import (
	"embed"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/plushgen"
	"github.com/gobuffalo/plush/v4"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/pop/v6/genny/config"
)

var AvailableDialects = pop.AvailableDialects

//go:embed templates/*
var templates embed.FS

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}

	if err := opts.Validate(); err != nil {
		return gg, err
	}

	sub, err := fs.Sub(templates, "templates")
	if err != nil {
		return gg, err
	}

	g := genny.New()
	if err := g.FS(sub); err != nil {
		return gg, err
	}

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())

	gg.Add(g)

	g, err = config.New(&config.Options{
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
