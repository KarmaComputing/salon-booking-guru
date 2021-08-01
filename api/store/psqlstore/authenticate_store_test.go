package psqlstore

import (
	"fmt"
	"reflect"
	"salon-booking-guru/store/model"
	"testing"
)

func TestAuthenticateCredentials(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	credentials := model.Credentials{
		Email:    "admin@example.com",
		Password: "password",
	}

	authenticateResponseGet, err := s.Authenticate().AuthenticateCredentials(credentials)
	if err != nil {
		t.Fatal(err)
	}

	authenticateResponse := model.AuthenticateResponse{
		AccountInfo: model.AccountInfo{
			Email:       "admin@example.com",
			FirstName:   "Adam",
			LastName:    "Appleby",
			RoleName:    "Administrator",
			Permissions: []string{"Administrator"},
		},
		Token: model.Token{
			AccountId: 1,
			Token:     "",
		},
	}

	authenticateResponseGet.Token.Token = ""

	if !reflect.DeepEqual(authenticateResponse, authenticateResponseGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", authenticateResponse, authenticateResponseGet))
	}
}
