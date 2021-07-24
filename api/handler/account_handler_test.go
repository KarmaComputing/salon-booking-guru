package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"salon-booking-guru/store/model"
	"salon-booking-guru/store/psqlstore"
	"testing"

	"github.com/gorilla/mux"
)

func TestAccountGetAll(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/account", nil)
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
}

func TestAccountGet(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/account/1", nil)
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
