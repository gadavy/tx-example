package ports

import "time"

type Clock interface {
	Now() time.Time
}

type TokenGenerator interface {
	Generate() (string, error)
	Verify(token string) bool
}

type BCryptGenerator interface {
	Generate(password string) (string, error)
	Compare(hashedPassword, password string) error
}
