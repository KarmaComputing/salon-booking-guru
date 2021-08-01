package psqlstore

import (
	"fmt"
	"salon-booking-guru/store/model"
	"testing"
)

func TestAuthorizeToken(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	credentials := model.Credentials{
		Email:    "permissiontest@example.com",
		Password: "password",
	}

	authenticateResponse, err := s.Authenticate().AuthenticateCredentials(credentials)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Authorize().AuthorizeToken(
		fmt.Sprintf("Bearer %s", authenticateResponse.Token.Token),
		[]string{},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Authorize().AuthorizeToken(
		fmt.Sprintf("Bearer %s", authenticateResponse.Token.Token),
		[]string{"canPermissionTest"},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Authorize().AuthorizeToken(
		fmt.Sprintf("Bearer %s", authenticateResponse.Token.Token),
		[]string{"thisPermissionDoesntExist"},
	)
	if err == nil {
		t.Fatal(err)
	}
}
