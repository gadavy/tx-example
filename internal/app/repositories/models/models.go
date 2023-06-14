package models

import (
	"time"

	"github.com/gadavy/tx-example/internal/app/domain"
)

type User struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
}

func UserFromDomain(value *domain.User) *User {
	return &User{
		ID:        value.ID,
		CreatedAt: value.CreatedAt,
		Email:     value.Email,
		Password:  value.Password,
	}
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		Email:     u.Email,
		Password:  u.Password,
	}
}

type Session struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserID    string    `db:"user_id"`
	Token     string    `db:"token"`
}

func SessionFromDomain(value *domain.Session) *Session {
	return &Session{
		ID:        value.ID,
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
		UserID:    value.UserID,
		Token:     value.Token,
	}
}

func (s *Session) ToDomain() *domain.Session {
	return &domain.Session{
		ID:        s.ID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		UserID:    s.UserID,
		Token:     s.Token,
	}
}
