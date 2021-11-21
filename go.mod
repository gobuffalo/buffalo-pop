module github.com/gobuffalo/buffalo-pop/v3

go 1.16

require (
	github.com/fatih/color v1.13.0
	github.com/gobuffalo/buffalo v0.17.5
	github.com/gobuffalo/events v1.4.2
	github.com/gobuffalo/flect v0.2.4
	github.com/gobuffalo/genny/v2 v2.0.8
	github.com/gobuffalo/httptest v1.5.1
	github.com/gobuffalo/packr/v2 v2.8.2
	github.com/gobuffalo/plush/v4 v4.1.9
	github.com/gobuffalo/pop/v5 v5.3.4
	github.com/gofrs/uuid v4.1.0+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/gobuffalo/buffalo v0.17.5 => github.com/fasmat/buffalo v0.16.15-0.20211121153216-deb763ccd343
	github.com/gobuffalo/pop/v6 v6.0.0 => github.com/fasmat/pop/v6 v6.0.0-20211121152713-e467c63e98c0
)
