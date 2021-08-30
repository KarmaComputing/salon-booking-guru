package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"salon-booking-guru/store/model"
	"salon-booking-guru/store/psqlstore"
	"testing"

	"github.com/gorilla/mux"
)

func TestAvailabilityGetAllByAccountId(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/account/3/availability", nil)
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
	var availabilities []model.Availability
	json.Unmarshal(bodyBytes, &availabilities)

	if len(availabilities) != 3 {
		t.Fatal(errors.New("Number of availabilities returned is invalid"))
	}
}

/* func TestAvailabilityGet(t *testing.T) {
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

func TestAvailabilityCreate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Availability{
		Name: "Test Availability",
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
	var originalAvailability model.Availability
	json.Unmarshal(bodyBytes, &originalAvailability)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/qualification/%d", originalAvailability.Id),
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
	var responseAvailability model.Availability
	json.Unmarshal(bodyBytes, &responseAvailability)

	if !reflect.DeepEqual(responseAvailability, originalAvailability) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseAvailability,
				originalAvailability,
			),
		)
	}
}

func TestAvailabilityCreateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Availability{
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

func TestAvailabilityUpdate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Availability{
		Id:   2,
		Name: "Updated Availability",
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
	var originalAvailability model.Availability
	json.Unmarshal(bodyBytes, &originalAvailability)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/qualification/%d", originalAvailability.Id),
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
	var responseAvailability model.Availability
	json.Unmarshal(bodyBytes, &responseAvailability)

	if !reflect.DeepEqual(responseAvailability, originalAvailability) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseAvailability,
				originalAvailability,
			),
		)
	}
}

func TestAvailabilityUpdateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	qualification := model.Availability{
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

func TestAvailabilityDelete(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest(
		"DELETE",
		"/v1/qualification/1",
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
} */
