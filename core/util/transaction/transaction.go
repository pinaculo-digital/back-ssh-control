package transaction

import (
	"context"
	"errors"
	"go_service/core/server/shared"

	"github.com/jmoiron/sqlx"
)

func RunInTx(fn func(tx *sqlx.Tx) error) error {
	tx, err := shared.GetDB().BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
