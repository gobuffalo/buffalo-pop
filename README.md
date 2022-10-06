<p align="center"><img src="https://github.com/gobuffalo/buffalo/blob/main/logo.svg" width="360"></p>

<p align="center">
<a href="https://godoc.org/github.com/gobuffalo/buffalo-pop"><img src="https://godoc.org/github.com/gobuffalo/buffalo-pop?status.svg" alt="GoDoc" /></a>
<a href="https://goreportcard.com/report/github.com/gobuffalo/buffalo-pop"><img src="https://goreportcard.com/badge/github.com/gobuffalo/buffalo-pop" alt="Go Report Card" /></a>
</p>

# github.com/gobuffalo/buffalo-pop

This is the home for all things that combine [Buffalo](https://github.com/gobuffalo/buffalo) and [Pop](https://github.com/gobuffalo/pop).

## Installation

```bash
go install github.com/gobuffalo/buffalo-pop/v3@latest
```

Or with SQLite 3 support:

```bash
go get -tags sqlite -v github.com/gobuffalo/buffalo-pop/v3
```

## Transaction Middleware

The `popmw.Transaction` will wrap each request inside of a new database transaction and automatically commit, or rollback, based on whether or not an error was returned from an upstream `buffalo.Handler` or `buffalo.MiddlewareFunc`.

### Usage

First you need to add the middleware to your application giving it access to your `*pop.Connection`, typically found at `models.DB`.

```go
import "github.com/gobuffalo/buffalo-pop/v3/pop/popmw"

func App() *buffalo.App {
  // ...
  app.Use(popmw.Transaction(models.DB))
  // ...
}
```

Once added to the middleware stack for your application you can then use this transaction in upstream middleware or handlers.

```go
func MyHandler(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }
}
```

**WARNING: DO NOT OPEN MULTIPLE TRANSACTIONS WITHIN EACH OTHER** - doing so will cause many, many, many problems.
