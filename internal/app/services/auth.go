package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gadavy/tx-example/internal/app/domain"
	"github.com/gadavy/tx-example/internal/app/ports"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"
)

const _sessionMaxAge = 8760 * time.Hour

type AuthService struct {
	logger tracelog.Logger
	clock  ports.Clock

	usersRepo    ports.UsersRepo
	sessionsRepo ports.SessionsRepo
	historyRepo  ports.LoginHistoryRepo

	tokenGen  ports.TokenGenerator
	bcryptGen ports.BCryptGenerator
}

func NewAuthService(
	logger tracelog.Logger,
	clock ports.Clock,
	usersRepo ports.UsersRepo,
	sessionsRepo ports.SessionsRepo,
	historyRepo ports.LoginHistoryRepo,
	tokenGen ports.TokenGenerator,
	bcryptGen ports.BCryptGenerator,
) *AuthService {
	return &AuthService{
		logger:       logger,
		clock:        clock,
		usersRepo:    usersRepo,
		sessionsRepo: sessionsRepo,
		historyRepo:  historyRepo,
		tokenGen:     tokenGen,
		bcryptGen:    bcryptGen,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, token string) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	if !s.tokenGen.Verify(token) {
		return nil, domain.ErrInvalidToken
	}

	session, err := s.sessionsRepo.FindByToken(ctx, token)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("find session: %w", err)
	}

	if errors.Is(err, domain.ErrNotFound) || session.IsExpired(s.clock.Now(), _sessionMaxAge) {
		return nil, domain.ErrTokenNotFound
	}

	if err = s.sessionsRepo.Update(ctx, session.ID); err != nil {
		return nil, fmt.Errorf("update session: %w", err)
	}

	return session, nil
}

func (s *AuthService) SignIn(ctx context.Context, params *domain.SignInParams) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	user, err := s.usersRepo.FindByEmail(ctx, params.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrInvalidLoginOrPassword
		}

		return nil, fmt.Errorf("find by email: %w", err)
	}

	if err := s.bcryptGen.Compare(user.Password, params.Password); err != nil {
		return nil, domain.ErrInvalidLoginOrPassword
	}

	var session *domain.Session

	if err := s.usersRepo.RunTxx(ctx, func(ctx context.Context) error {
		token, err := s.tokenGen.Generate()
		if err != nil {
			return fmt.Errorf("generate token: %w", err)
		}

		session, err = s.sessionsRepo.Create(ctx, user.NewSession(token))
		if err != nil {
			return fmt.Errorf("create session: %w", err)
		}

		if err := s.historyRepo.Create(ctx, user.ID, params.IP, params.Useragent); err != nil {
			return fmt.Errorf("create login history: %w", err)
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("tx: %w", err)
	}

	return session, nil
}

func (s *AuthService) SignUp(ctx context.Context, params *domain.SignUpParams) error {
	defer tracing.ChildSpan(&ctx).Finish()

	exists, err := s.usersRepo.FindByEmail(ctx, params.Email)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("find by email: %w", err)
	}

	if exists != nil {
		return domain.ErrEmailAlreadyExists
	}

	password, err := s.bcryptGen.Generate(params.Password)
	if err != nil {
		return fmt.Errorf("generate password hash failed: %w", err)
	}

	_, err = s.usersRepo.Create(ctx, &domain.User{
		Email:    params.Email,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	return nil
}

func (s *AuthService) SignOut(ctx context.Context, session *domain.Session) error {
	defer tracing.ChildSpan(&ctx).Finish()

	if err := s.sessionsRepo.RunTxx(ctx, func(ctx context.Context) error {
		session, err := s.sessionsRepo.FindByToken(ctx, session.Token)
		if err != nil {
			return fmt.Errorf("find session by token: %w", err)
		}

		if err := s.sessionsRepo.Delete(ctx, session.ID); err != nil {
			return fmt.Errorf("delete session: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("tx: %w", err)
	}

	return nil
}
