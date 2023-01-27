package popmw

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

type widget struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const mig = `CREATE TABLE "widgets" (
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL,
	"id" TEXT PRIMARY KEY
  );`

func tx(fn func(tx *pop.Connection)) error {
	pop.Debug = true
	defer func() { pop.Debug = false }()

	d, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	path := filepath.Join(d, "pt_test.sqlite")
	defer os.RemoveAll(path)

	db, err := pop.NewConnection(&pop.ConnectionDetails{
		Dialect: "sqlite",
		URL:     path,
	})
	if err != nil {
		return err
	}

	if err := db.Dialect.CreateDB(); err != nil {
		return err
	}
	if err := db.Open(); err != nil {
		return err
	}
	if err := db.RawQuery(mig).Exec(); err != nil {
		return err
	}

	fn(db)

	return nil
}

func app(db *pop.Connection) *buffalo.App {
	app := buffalo.New(buffalo.Options{})
	app.Use(Transaction(db))

	// all handlers do the same database job but the return differently
	app.GET("/success-201", func(c buffalo.Context) error {
		w := &widget{}
		tx := c.Value("tx").(*pop.Connection)
		if err := tx.Create(w); err != nil {
			return err
		}
		return c.Render(201, nil) // 201 created
	})

	app.GET("/success-301", func(c buffalo.Context) error {
		w := &widget{}
		tx := c.Value("tx").(*pop.Connection)
		if err := tx.Create(w); err != nil {
			return err
		}
		return c.Render(301, nil) // 301 moved permanently
	})

	app.GET("/success-nil", func(c buffalo.Context) error {
		w := &widget{}
		tx := c.Value("tx").(*pop.Connection)
		if err := tx.Create(w); err != nil {
			return err
		}
		return nil // will become 200 ok
	})

	app.GET("/error-409", func(c buffalo.Context) error {
		w := &widget{}
		tx := c.Value("tx").(*pop.Connection)
		if err := tx.Create(w); err != nil {
			return err
		}
		return c.Render(409, nil) // 409 conflict
	})

	app.GET("/error-500", func(c buffalo.Context) error {
		w := &widget{}
		tx := c.Value("tx").(*pop.Connection)
		if err := tx.Create(w); err != nil {
			return err
		}
		return c.Render(500, nil) // 500 internal server error
	})

	app.GET("/error", func(c buffalo.Context) error {
		w := &widget{}
		tx := c.Value("tx").(*pop.Connection)
		if err := tx.Create(w); err != nil {
			return err
		}
		return fmt.Errorf("boom") // will become 500
	})

	return app
}

func Test_PopTransaction_Success(t *testing.T) {
	tests := []struct {
		path          string
		status        int
		success       bool
		expectedCount int
	}{
		{"success-201", 201, true, 1},
		{"success-301", 301, true, 1},
		{"success-nil", 200, true, 1},
		{"error-409", 409, true, 0},
		{"error-500", 500, true, 0},
		{"error", 500, true, 0},
	}
	for _, tc := range tests {
		t.Run(tc.path, func(t *testing.T) {
			r := require.New(t)
			err := tx(func(db *pop.Connection) {
				w := httptest.New(app(db))
				res := w.HTML("/" + tc.path).Get()
				r.Equal(tc.status, res.Code)

				count, err := db.Count("widgets")
				r.NoError(err)
				r.Equal(tc.expectedCount, count)
			})
			r.NoError(err)
		})
	}
}
