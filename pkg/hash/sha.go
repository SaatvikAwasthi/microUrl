package hash

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"hash"
)

type ShaOptions func() bool

func WithPadding() ShaOptions {
	return func() bool {
		return true
	}
}

type SHA struct {
	hash       hash.Hash
	addPadding bool
}

func New(options ...ShaOptions) *SHA {
	padding := false
	for _, option := range options {
		padding = option()

	}
	return &SHA{
		hash:       sha256.New(),
		addPadding: padding,
	}
}

func (s *SHA) Hash(data string) (string, error) {
	if _, err := s.hash.Write([]byte(data)); err != nil {
		return "", fmt.Errorf("failed to hash data. error: %w", err)
	}
	return string(s.insertRandomPadding()), nil
}

func (s *SHA) insertRandomPadding() []byte {
	id := uuid.New()
	if s.addPadding {
		return s.hash.Sum([]byte(id.String()))
	}
	return s.hash.Sum(nil)
}
