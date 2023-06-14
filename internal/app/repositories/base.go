package repositories

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var _transactionKey = struct{}{} //nolint:gochecknoglobals // best practice.

type DB interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PrepareNamedContext(ctx context.Context, query string) (stmt *sqlx.NamedStmt, err error)
}

type BaseRepo struct {
	__db *sqlx.DB //nolint:revive,stylecheck // for visibility illegal use
}

func NewBaseRepo(db *sqlx.DB) BaseRepo {
	return BaseRepo{__db: db}
}

func (r *BaseRepo) RunTxx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := r.__db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer r.rollback(tx)

	if err := fn(context.WithValue(ctx, _transactionKey, tx)); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err //nolint:wrapcheck // need clean error
	}

	return nil
}

func (r *BaseRepo) db(ctx context.Context) DB {
	if tx, ok := ctx.Value(_transactionKey).(*sqlx.Tx); ok {
		return tx
	}

	return r.__db
}

func (r *BaseRepo) rollback(tx *sqlx.Tx) {
	_ = tx.Rollback()
}
