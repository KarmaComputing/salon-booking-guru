package psqlstore

import (
	"testing"
)

var s, err = OpenTest()

func TestAccountGetAll(t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}

	_, err := s.Account().GetAll()
	if err != nil {
		t.Fatal(err)
	}
}
