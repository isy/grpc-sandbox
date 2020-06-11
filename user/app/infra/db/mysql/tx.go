package mysql

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func InTx(db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("Rollback error was occurred: %+v", err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("Rollback error was occurred: %+v", err)
		}
		return err
	}

	return nil
}
