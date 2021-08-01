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

func authorizeAsAdmin(t *testing.T, roleName string, req *http.Request) {
	credentials := model.Credentials{
		Email:    "admin@example.com",
		Password: "password",
	}

	authenticateResponse, err := s.Authenticate().AuthenticateCredentials(
		credentials,
	)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", authenticateResponse.Token.Token),
	)
}

func TestAccountGetAll(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/account", nil)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusOK,
		)
	}
}

func TestAccountGet(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/account/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusOK,
		)
	}
}

func TestAccountCreate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	mobileNumber := "07123456789"
	account := model.Account{
		RoleId:       1,
		FirstName:    "Test",
		LastName:     "User",
		Email:        "test@example.com",
		Password:     "password",
		MobileNumber: &mobileNumber,
	}

	accountJson, err := json.Marshal(account)

	req, err := http.NewRequest(
		"POST",
		"/v1/account",
		bytes.NewBuffer(accountJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

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
	var originalAccount model.Account
	json.Unmarshal(bodyBytes, &originalAccount)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/account/%d", originalAccount.Id),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr = httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusOK,
		)
	}

	bodyBytes, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	var responseAccount model.Account
	json.Unmarshal(bodyBytes, &responseAccount)

	responseAccount.Password = ""
	originalAccount.Password = ""

	if !reflect.DeepEqual(responseAccount, originalAccount) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseAccount,
				originalAccount,
			),
		)
	}
}

func TestAccountCreateInvalidEmail(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	mobileNumber := "07123456789"
	account := model.Account{
		RoleId:       1,
		FirstName:    "Test",
		LastName:     "User",
		Email:        "testexample.com",
		Password:     "password",
		MobileNumber: &mobileNumber,
	}

	accountJson, err := json.Marshal(account)

	req, err := http.NewRequest(
		"POST",
		"/v1/account",
		bytes.NewBuffer(accountJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestAccountUpdate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	mobileNumber := "07123456789"
	account := model.Account{
		RoleId:       2,
		FirstName:    "TestUpdated",
		LastName:     "UserUpdated",
		Email:        "test@example.com",
		Password:     "password",
		MobileNumber: &mobileNumber,
	}

	accountJson, err := json.Marshal(account)

	req, err := http.NewRequest(
		"POST",
		"/v1/account",
		bytes.NewBuffer(accountJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

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
	var originalAccount model.Account
	json.Unmarshal(bodyBytes, &originalAccount)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/account/%d", originalAccount.Id),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr = httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusOK,
		)
	}

	bodyBytes, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	var responseAccount model.Account
	json.Unmarshal(bodyBytes, &responseAccount)

	responseAccount.Password = ""
	originalAccount.Password = ""

	if !reflect.DeepEqual(responseAccount, originalAccount) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseAccount,
				originalAccount,
			),
		)
	}
}

func TestAccountUpdateInvalidEmail(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	mobileNumber := "07123456789"
	account := model.Account{
		RoleId:       1,
		FirstName:    "TestUpdated",
		LastName:     "UserUpdated",
		Email:        "testexample.com",
		Password:     "password",
		MobileNumber: &mobileNumber,
	}

	accountJson, err := json.Marshal(account)

	req, err := http.NewRequest(
		"POST",
		"/v1/account",
		bytes.NewBuffer(accountJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestAccountDelete(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest(
		"DELETE",
		"/v1/account/5",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, "Administrator", req)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Handler returned wrong status code: got %v, want %v",
			status,
			http.StatusOK,
		)
	}
}
