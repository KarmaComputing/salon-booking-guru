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
		Email:    "permissiontest@example.com",
		Password: "password",
	}

	authenticateResponseGet, err := s.Authenticate().AuthenticateCredentials(credentials)
	if err != nil {
		t.Fatal(err)
	}

	authenticateResponse := model.AuthenticateResponse{
		AccountInfo: model.AccountInfo{
			Email:       "permissiontest@example.com",
			FirstName:   "Edgar",
			LastName:    "Evans",
			RoleName:    "PermissionTest",
			Permissions: []string{"canPermissionTest"},
		},
		Token: model.Token{
			AccountId: 5,
			Token:     "",
		},
	}

	authenticateResponseGet.Token.Token = ""

	if !reflect.DeepEqual(authenticateResponse, authenticateResponseGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", authenticateResponse, authenticateResponseGet))
	}
}
