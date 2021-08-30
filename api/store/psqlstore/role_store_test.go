package psqlstore

import (
	"errors"
	"testing"
)

func TestRoleGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	accounts, err := s.Role().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) != 4 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}
