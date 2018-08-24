package newapp

import "github.com/gobuffalo/buffalo/meta"

type Options struct {
	App     meta.App
	Dialect string
	Prefix  string
}
