package crypto

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     = bcrypt.MinCost
	MaxCost     = bcrypt.MaxCost
	DefaultCost = bcrypt.DefaultCost
)

type BCrypt struct {
	cost int
}

func NewBCrypt(cost int) *BCrypt {
	return &BCrypt{
		cost: cost,
	}
}

func (b *BCrypt) Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", fmt.Errorf("generate: %w", err)
	}

	return hex.EncodeToString(hash), nil
}

func (b *BCrypt) Compare(hashedPassword, password string) error {
	hash, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return fmt.Errorf("decode hashed password: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return fmt.Errorf("compare: %w", err)
	}

	return nil
}
