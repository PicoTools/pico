package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/PicoTools/pico/internal/ent"
	"github.com/go-faster/errors"
	"modernc.org/sqlite"

	"github.com/go-faster/sdk/zctx"
	"go.uber.org/zap"
)

// sqlite3 config
type ConfigSqlite struct {
	Path string `json:"path" validate:"required"`
}

// Init initialize sqlite3 connection
func (c *ConfigSqlite) Init(ctx context.Context) (*ent.Client, error) {
	lg := zctx.From(ctx).Named("sqlite")

	var d strings.Builder
	d.WriteString("file:")
	d.WriteString(c.Path)
	d.WriteString("?cache=shared&_fk=1")

	lg.Debug("connection dsn", zap.String("dsn", d.String()))

	db, err := sql.Open("sqlite3", d.String())
	if err != nil {
		return nil, err
	}
	// avoid "database is locked" errors
	db.SetMaxOpenConns(1)

	// create client
	drv := entsql.OpenDB(dialect.SQLite, db)
	client := ent.NewClient(ent.Driver(drv))
	lg.Debug("connected to database")

	// migrations
	if err = client.Schema.Create(ctx); err != nil {
		_ = client.Close()
		return nil, err
	}
	lg.Debug("schema migrated")

	return client, nil
}

type sqlite3Driver struct {
	*sqlite.Driver
}

// Open create connection to database
func (d sqlite3Driver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	// enable PRAGMAs
	if _, err = c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		_ = conn.Close()
		return nil, errors.Wrap(err, "failed to enable foreign keys")
	}
	if _, err = c.Exec("PRAGMA journal_mode = WAL;", nil); err != nil {
		_ = conn.Close()
		return nil, errors.Wrap(err, "failed to enable WAL mode")
	}
	if _, err = c.Exec("PRAGMA synchronous=NORMAL;", nil); err != nil {
		_ = conn.Close()
		return nil, errors.Wrap(err, "failed to enable normal synchronous")
	}
	return conn, nil
}

func init() {
	// register sqlite3 driver for using with ent
	sql.Register("sqlite3", sqlite3Driver{Driver: &sqlite.Driver{}})
}
