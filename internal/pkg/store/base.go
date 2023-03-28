package store

import (
	"github.com/go-pg/pg"
)

type Database struct {
	*pg.DB
}

func (db *Database) Error() string {
	//TODO implement me
	panic("implement me")
}

func (db *Database) Heartbeat() error {
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return err
	}

	return nil
}
