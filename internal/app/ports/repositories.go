package ports

import (
	"context"

	"github.com/gadavy/tx-example/internal/app/domain"
)

type TxRepo interface {
	RunTxx(ctx context.Context, fn func(ctx context.Context) error) error
}

type UsersRepo interface {
	TxRepo

	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type SessionsRepo interface {
	TxRepo

	Create(ctx context.Context, session *domain.Session) (*domain.Session, error)
	FindByID(ctx context.Context, sessionID string) (*domain.Session, error)
	FindByToken(ctx context.Context, token string) (*domain.Session, error)
	Update(ctx context.Context, sessionID string) error
	Delete(ctx context.Context, sessionID string) error
}

type LoginHistoryRepo interface {
	TxRepo

	Create(ctx context.Context, userID, idAddr, userAgent string) error
}
