package hash

import (
	"crypto/sha1"
	"fmt"
)

type SHA1 struct {
	salt string
}

func NewSHA1(salt string) *SHA1 {
	return &SHA1{salt: salt}
}

func (s *SHA1) HashPassword(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt))), nil
}
