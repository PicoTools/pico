package db

import (
	"context"

	"github.com/PicoTools/pico/internal/ent"

	"github.com/go-faster/errors"
)

// Init initialize ent client with underly DB engine
func (c Config) Init(ctx context.Context) (*ent.Client, error) {
	if c.Sqlite != nil {
		db, err := c.Sqlite.Init(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "init sqlite")
		}
		return db, nil
	}

	// we must not get here
	return nil, errors.New("no db selected for processing")
}
