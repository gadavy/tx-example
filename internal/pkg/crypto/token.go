package crypto

import (
	"bytes"
	cryptoRand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

const (
	KeySize  = 32
	SignSize = sha256.Size
)

var (
	ErrBadToken    = errors.New("bad token")
	ErrEmptySecret = errors.New("empty secret")
)

type TokenGenerator struct {
	rnd    io.Reader
	secret []byte
}

func NewTokenGenerator(secret string) *TokenGenerator {
	return &TokenGenerator{
		secret: []byte(secret),
		rnd:    cryptoRand.Reader,
	}
}

func (g *TokenGenerator) Generate() (string, error) {
	return g.generateToken(g.rnd)
}

func (g *TokenGenerator) Verify(token string) bool {
	value, err := hex.DecodeString(token)
	if err != nil {
		return false
	}

	if !tokenLenIsValid(value) {
		return false
	}

	payload, sign := value[:len(value)-SignSize], value[len(value)-SignSize:]

	return bytes.Equal(g.sign(payload), sign)
}

func (g *TokenGenerator) generateToken(r io.Reader) (string, error) {
	buf := make([]byte, KeySize, KeySize+SignSize)

	if _, err := r.Read(buf[:KeySize]); err != nil {
		return "", fmt.Errorf("read rand bytes: %w", err)
	}

	buf = append(buf, g.sign(buf)...)

	return hex.EncodeToString(buf), nil
}

func (g *TokenGenerator) sign(message []byte) []byte {
	h := sha256.New()
	_, _ = h.Write(message)  // error always nil.
	_, _ = h.Write(g.secret) // error always nil.

	return h.Sum(nil)
}

func tokenLenIsValid(b []byte) bool {
	return len(b) == KeySize+SignSize
}
