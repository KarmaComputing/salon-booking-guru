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

func TestQualificationGetAll(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/qualification", nil)
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

func TestQualificationGet(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/qualification/1", nil)
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

func TestQualificationCreate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Qualification{
		Name: "Test Qualification",
	}

	qualificationJson, err := json.Marshal(qualification)

	req, err := http.NewRequest(
		"POST",
		"/v1/qualification",
		bytes.NewBuffer(qualificationJson),
	)
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

	bodyBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	var originalQualification model.Qualification
	json.Unmarshal(bodyBytes, &originalQualification)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/qualification/%d", originalQualification.Id),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, req)

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
	var responseQualification model.Qualification
	json.Unmarshal(bodyBytes, &responseQualification)

	if !reflect.DeepEqual(responseQualification, originalQualification) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseQualification,
				originalQualification,
			),
		)
	}
}

func TestQualificationCreateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Qualification{
		Name: "",
	}

	qualificationJson, err := json.Marshal(qualification)

	req, err := http.NewRequest(
		"POST",
		"/v1/qualification",
		bytes.NewBuffer(qualificationJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, req)

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

func TestQualificationUpdate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Qualification{
		Id:   2,
		Name: "Updated Qualification",
	}

	qualificationJson, err := json.Marshal(qualification)

	req, err := http.NewRequest(
		"PUT",
		"/v1/qualification",
		bytes.NewBuffer(qualificationJson),
	)
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

	bodyBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	var originalQualification model.Qualification
	json.Unmarshal(bodyBytes, &originalQualification)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/qualification/%d", originalQualification.Id),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, req)

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
	var responseQualification model.Qualification
	json.Unmarshal(bodyBytes, &responseQualification)

	if !reflect.DeepEqual(responseQualification, originalQualification) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseQualification,
				originalQualification,
			),
		)
	}
}

func TestQualificationUpdateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Qualification{
		Id:   2,
		Name: "",
	}

	qualificationJson, err := json.Marshal(qualification)

	req, err := http.NewRequest(
		"PUT",
		"/v1/qualification",
		bytes.NewBuffer(qualificationJson),
	)
	if err != nil {
		t.Fatal(err)
	}

	authorizeAsAdmin(t, req)

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

func TestQualificationDelete(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest(
		"DELETE",
		"/v1/qualification/2",
		nil,
	)
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
