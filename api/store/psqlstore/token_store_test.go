package psqlstore

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	token, err := s.Token().Generate(1)
	if err != nil {
		t.Fatal(err)
	}

	if token.AccountId != 1 {
		t.Fatal("Account id is incorrect")
	}

	if len(token.Token) != 88 {
		t.Fatal("Generated token is the incorrect length")
	}
}
