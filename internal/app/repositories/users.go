package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gadavy/tx-example/internal/app/domain"
	"github.com/gadavy/tx-example/internal/app/ports"
	"github.com/gadavy/tx-example/internal/app/repositories/models"
	"github.com/jmoiron/sqlx"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"
)

var _ ports.UsersRepo = &UsersRepo{}

type UsersRepo struct {
	BaseRepo
	logger tracelog.Logger
}

func NewUsersRepo(logger tracelog.Logger, db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		BaseRepo: NewBaseRepo(db),
		logger:   logger,
	}
}

func (r *UsersRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `
	INSERT INTO users (
		email,
		password
	) VALUES (
		:email,
		:password
	) RETURNING
		id,
		created_at,
		updated_at,
		email,
		password`

	stmt, err := r.db(ctx).PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("prepare stmt: %w", err)
	}

	var dest models.User

	if err := stmt.GetContext(ctx, &dest, models.UserFromDomain(user)); err != nil {
		return nil, fmt.Errorf("get context: %w", err)
	}

	return dest.ToDomain(), nil
}

func (r *UsersRepo) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `
	SELECT 
		id,
		created_at,
		updated_at,
		email,
		password
	FROM 
		users 
	WHERE 
		id = $1`

	var dest models.User

	if err := r.db(ctx).GetContext(ctx, &dest, query, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("get context: %w", err)
	}

	return dest.ToDomain(), nil
}

func (r *UsersRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	const query = `
	SELECT 
		id,
		created_at,
		updated_at,
		email,
		password
	FROM 
		users
	WHERE 
		email = $1`

	var dest models.User

	if err := r.db(ctx).GetContext(ctx, &dest, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("get context: %w", err)
	}

	return dest.ToDomain(), nil
}
