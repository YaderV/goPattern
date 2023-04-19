package strategy_test

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/YaderV/goPattern/strategy"
)

func TestStrategy(t *testing.T) {
	password := "qwerty"

	shaAlgorith := &strategy.SHA{}
	p := &strategy.Password{
		Password:  password,
		Algorithm: shaAlgorith,
	}

	shaHash := sha256.Sum256([]byte(password))
	sha := string(shaHash[:])

	if p.Hash() != sha {
		t.Fatal("The sha hash is not the same")
	}

	md5Algorith := &strategy.MD5{}
	p.SetAlgorithm(md5Algorith)
	md5Hash := md5.Sum([]byte(password))
	md5 := hex.EncodeToString(md5Hash[:])

	if p.Hash() != md5 {
		t.Fatal("The md5 hash is not the same")
	}

}
