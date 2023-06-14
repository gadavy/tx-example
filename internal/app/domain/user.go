package domain

import (
	"context"
	"time"
)

var _sessionKey = struct{}{} //nolint:gochecknoglobals // best practice.

type SignInParams struct {
	Email     string
	Password  string
	IP        string
	Useragent string
}

type SignUpParams struct {
	Email     string
	Password  string
	IP        string
	Useragent string
}

type User struct {
	ID        string
	CreatedAt time.Time
	Email     string
	Password  string
	IsBlocked bool
}

func (u *User) NewSession(token string) *Session {
	return &Session{
		UserID: u.ID,
		Token:  token,
	}
}

type Session struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	Token     string
}

func (s *Session) IsExpired(now time.Time, ttl time.Duration) bool {
	return now.After(s.UpdatedAt.Add(ttl))
}

func SessionToContext(ctx context.Context, user *Session) context.Context {
	return context.WithValue(ctx, _sessionKey, user)
}

func SessionFromContext(ctx context.Context) (*Session, bool) {
	if val, ok := ctx.Value(_sessionKey).(*Session); ok {
		return val, true
	}

	return nil, false
}
