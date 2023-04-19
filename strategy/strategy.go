package strategy

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// HashAlgorithm ...
type HashAlgorithm interface {
	Hash(p string) string
}

// Password ...
type Password struct {
	Password  string
	Algorithm HashAlgorithm
}

// Hash ...
func (p Password) Hash() string {
	return p.Algorithm.Hash(p.Password)
}

// SetAlgorithm ..
func (p *Password) SetAlgorithm(h HashAlgorithm) {
	p.Algorithm = h
}

// SHA ...
type SHA struct{}

// Hash ...
func (s *SHA) Hash(p string) string {
	hash := sha256.Sum256([]byte(p))
	return string(hash[:])
}

// MD5 ...
type MD5 struct{}

// Hash ...
func (s *MD5) Hash(p string) string {
	hash := md5.Sum([]byte(p))
	return hex.EncodeToString(hash[:])
}
