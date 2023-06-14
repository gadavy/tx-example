package ports

import (
	"context"

	"github.com/gadavy/tx-example/internal/app/domain"
)

type AuthService interface {
	SignIn(ctx context.Context, params *domain.SignInParams) (*domain.Session, error)
	SignUp(ctx context.Context, params *domain.SignUpParams) error
	SignOut(ctx context.Context, session *domain.Session) error
	Authenticate(ctx context.Context, token string) (*domain.Session, error)
}
