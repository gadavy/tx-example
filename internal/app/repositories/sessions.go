package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gadavy/tx-example/internal/app/domain"
	"github.com/gadavy/tx-example/internal/app/ports"
	"github.com/gadavy/tx-example/internal/app/repositories/models"
	"github.com/jmoiron/sqlx"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"
)

var _ ports.SessionsRepo = &SessionsRepo{}

type SessionsRepo struct {
	BaseRepo
	logger tracelog.Logger
}

func NewSessionRepo(logger tracelog.Logger, db *sqlx.DB) *SessionsRepo {
	return &SessionsRepo{
		BaseRepo: NewBaseRepo(db),
		logger:   logger,
	}
}

func (r *SessionsRepo) Create(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `INSERT INTO users_session (
		user_id,
		token
	) VALUES (
		:user_id,
		:token
	) RETURNING id`

	stmt, err := r.db(ctx).PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("prepare stmt: %w", err)
	}

	if err := stmt.GetContext(ctx, &session.ID, models.SessionFromDomain(session)); err != nil {
		return nil, fmt.Errorf("get context: %w", err)
	}

	return session, nil
}

func (r *SessionsRepo) FindByID(ctx context.Context, sessionID string) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `SELECT 
		id,
		created_at,
		updated_at,
		user_id,
		token
    FROM
        users_session 
    WHERE id = $1`

	var dest models.Session

	if err := r.db(ctx).GetContext(ctx, &dest, query, sessionID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("get context: %w", err)
	}

	return dest.ToDomain(), nil
}

func (r *SessionsRepo) FindByToken(ctx context.Context, token string) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `SELECT 
		id,
		created_at,
		updated_at,
		user_id,
		token
    FROM
        users_session 
    WHERE token = $1`

	var dest models.Session

	if err := r.db(ctx).GetContext(ctx, &dest, query, token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("get context: %w", err)
	}

	return dest.ToDomain(), nil
}

func (r *SessionsRepo) Update(ctx context.Context, sessionID string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `UPDATE users_session SET updated_at = $1 WHERE id = $2`

	if _, err := r.db(ctx).ExecContext(ctx, query, time.Now(), sessionID); err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (r *SessionsRepo) Delete(ctx context.Context, sessionID string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `DELETE FROM users_session WHERE id = $1`

	if _, err := r.db(ctx).ExecContext(ctx, query, sessionID); err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}
