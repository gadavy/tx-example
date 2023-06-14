package repositories

import (
	"context"
	"fmt"

	"github.com/gadavy/tx-example/internal/app/ports"
	"github.com/jmoiron/sqlx"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"
)

var _ ports.LoginHistoryRepo = &LoginHistoryRepo{}

type LoginHistoryRepo struct {
	BaseRepo
	logger tracelog.Logger
}

func NewLoginHistoryRepo(logger tracelog.Logger, db *sqlx.DB) *LoginHistoryRepo {
	return &LoginHistoryRepo{
		BaseRepo: NewBaseRepo(db),
		logger:   logger,
	}
}

func (r *LoginHistoryRepo) Create(ctx context.Context, userID, idAddr, userAgent string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `
	INSERT INTO users_login_history (
		user_id,
		ip_addr,
		useragent
	) VALUES (
		$1,$2,$3
	)`

	if _, err := r.db(ctx).ExecContext(ctx, query, userID, idAddr, userAgent); err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}
