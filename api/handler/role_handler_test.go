package handler

import (
	"net/http"
	"net/http/httptest"
	"salon-booking-guru/store/psqlstore"
	"testing"

	"github.com/gorilla/mux"
)

func TestRoleGetAll(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/role", nil)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, req)

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
