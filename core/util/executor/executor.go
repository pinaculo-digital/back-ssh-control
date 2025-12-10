package executor

import (
	"context"
	"database/sql"
	"go_service/core/server/shared"

	"github.com/jmoiron/sqlx"
)

type Executor interface {
	// Métodos COM Context (para operações assíncronas/timeouts)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	// Métodos SEM Context (para operações síncronas simples)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error

	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error

	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)

	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryRowx(query string, args ...interface{}) *sqlx.Row

	StructScanContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	StructScan(dest interface{}, query string, args ...interface{}) error
}

// DBExecutor implementação concreta
type DBExecutor struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewDBExecutor(tx *sqlx.Tx) Executor {
	return &DBExecutor{db: shared.DB, tx: tx}
}

// Métodos COM Context
func (e *DBExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if e.tx != nil {
		return e.tx.ExecContext(ctx, query, args...)
	}
	return e.db.ExecContext(ctx, query, args...)
}

func (e *DBExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if e.tx != nil {
		return e.tx.QueryContext(ctx, query, args...)
	}
	return e.db.QueryContext(ctx, query, args...)
}

func (e *DBExecutor) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if e.tx != nil {
		return e.tx.QueryRowContext(ctx, query, args...)
	}
	return e.db.QueryRowContext(ctx, query, args...)
}

func (e *DBExecutor) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if e.tx != nil {
		return e.tx.PrepareContext(ctx, query)
	}
	return e.db.PrepareContext(ctx, query)
}

// Métodos SEM Context (usam context.Background())
func (e *DBExecutor) Exec(query string, args ...interface{}) (sql.Result, error) {
	return e.ExecContext(context.Background(), query, args...)
}

func (e *DBExecutor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return e.QueryContext(context.Background(), query, args...)
}

func (e *DBExecutor) QueryRow(query string, args ...interface{}) *sql.Row {
	return e.QueryRowContext(context.Background(), query, args...)
}

func (e *DBExecutor) Prepare(query string) (*sql.Stmt, error) {
	return e.PrepareContext(context.Background(), query)
}

func (e *DBExecutor) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if e.tx != nil {
		return e.tx.SelectContext(ctx, dest, query, args...)
	}
	return e.db.SelectContext(ctx, dest, query, args...)
}

func (e *DBExecutor) Select(dest interface{}, query string, args ...interface{}) error {
	return e.SelectContext(context.Background(), dest, query, args...)
}

func (e *DBExecutor) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if e.tx != nil {
		return e.tx.GetContext(ctx, dest, query, args...)
	}
	return e.db.GetContext(ctx, dest, query, args...)
}

func (e *DBExecutor) Get(dest interface{}, query string, args ...interface{}) error {
	return e.GetContext(context.Background(), dest, query, args...)
}

func (e *DBExecutor) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	if e.tx != nil {
		return e.tx.QueryxContext(ctx, query, args...)
	}
	return e.db.QueryxContext(ctx, query, args...)
}

func (e *DBExecutor) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return e.QueryxContext(context.Background(), query, args...)
}

func (e *DBExecutor) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	if e.tx != nil {
		return e.tx.QueryRowxContext(ctx, query, args...)
	}
	return e.db.QueryRowxContext(ctx, query, args...)
}

func (e *DBExecutor) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return e.QueryRowxContext(context.Background(), query, args...)
}

// StructScan em uma única linha (útil em loops)
func (e *DBExecutor) StructScanContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	row := e.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}

func (e *DBExecutor) StructScan(dest interface{}, query string, args ...interface{}) error {
	return e.StructScanContext(context.Background(), dest, query, args...)
}
