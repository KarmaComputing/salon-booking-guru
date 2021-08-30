package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"salon-booking-guru/store/model"
	"salon-booking-guru/store/psqlstore"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestAvailabilityGet(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/availability/1", nil)
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

func TestAvailabilityCreate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	availabilities := []model.Availability{
		model.Availability{
			AccountId: 1,
			StartDate: time.Date(2021, time.Month(5), 17, 9, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2021, time.Month(5), 18, 17, 0, 0, 0, time.UTC),
		},
		model.Availability{
			AccountId: 1,
			StartDate: time.Date(2021, time.Month(5), 18, 9, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2021, time.Month(5), 18, 17, 0, 0, 0, time.UTC),
		},
		model.Availability{
			AccountId: 1,
			StartDate: time.Date(2021, time.Month(5), 19, 9, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2021, time.Month(5), 19, 17, 0, 0, 0, time.UTC),
		},
	}

	availabilityJson, err := json.Marshal(availabilities)

	req, err := http.NewRequest(
		"POST",
		"/v1/availability",
		bytes.NewBuffer(availabilityJson),
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
	var originalAvailabilities []model.Availability
	json.Unmarshal(bodyBytes, &originalAvailabilities)

	req, err = http.NewRequest(
		"GET",
		"/v1/account/1/availability",
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
	var responseAvailabilities []model.Availability
	json.Unmarshal(bodyBytes, &responseAvailabilities)

	if len(responseAvailabilities) != 3 {
		t.Fatal(errors.New("Number of availabilities returned is invalid"))
	}

	for i := 0; i < len(responseAvailabilities); i++ {
		responseAvailabilities[i].Id = 0
	}

	if !reflect.DeepEqual(responseAvailabilities, originalAvailabilities) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseAvailabilities,
				originalAvailabilities,
			),
		)
	}
}

func TestAvailabilityUpdate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	availability := model.Availability{
		Id:        1,
		AccountId: 3,
		StartDate: time.Date(2021, time.Month(5), 17, 9, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2021, time.Month(5), 17, 17, 0, 0, 0, time.UTC),
	}

	availabilityJson, err := json.Marshal(availability)

	req, err := http.NewRequest(
		"PUT",
		"/v1/availability",
		bytes.NewBuffer(availabilityJson),
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
		fmt.Sprintf("/v1/availability/%d", originalAvailability.Id),
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

func TestAvailabilityDelete(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest(
		"DELETE",
		"/v1/availability/1",
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
