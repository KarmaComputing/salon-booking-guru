package psqlstore

import (
	"testing"
)

func TestAccountGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Account().GetAll()
	if err != nil {
		t.Fatal(err)
	}
}
