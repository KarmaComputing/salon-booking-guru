package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"salon-booking-guru/store/model"
	"salon-booking-guru/store/psqlstore"
	"testing"

	"github.com/gorilla/mux"
)

func TestAuthenticateCredentials(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	credentials := model.Credentials{
		Email:    "admin@example.com",
		Password: "password",
	}

	credentialsJson, err := json.Marshal(credentials)

	req, err := http.NewRequest(
		"POST",
		"/v1/authenticate",
		bytes.NewBuffer(credentialsJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusOK,
		)
	}

	bodyBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	var authenticateResponseResponse model.AuthenticateResponse
	json.Unmarshal(bodyBytes, &authenticateResponseResponse)
	authenticateResponseResponse.Token.Token = ""

	authenticateResponse := model.AuthenticateResponse{
		AccountInfo: model.AccountInfo{
			Email:     "admin@example.com",
			FirstName: "Adam",
			LastName:  "Appleby",
			RoleName:  "Administrator",
		},
		Token: model.Token{
			AccountId: 1,
			Token:     "",
		},
	}

	if !reflect.DeepEqual(authenticateResponseResponse, authenticateResponse) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				authenticateResponseResponse,
				authenticateResponse,
			),
		)
	}
}
