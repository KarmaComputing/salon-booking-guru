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

func TestProductGetAll(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/product", nil)
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

func TestProductGet(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/product/1", nil)
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

func TestProductCreate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	product := model.Product{
		ProductCategoryId: 2,
		Name:              "Product Test",
		Description:       "Product Test description.",
		Price:             1.99,
		Deposit:           0.99,
		Duration:          0.5,
	}

	productJson, err := json.Marshal(product)

	req, err := http.NewRequest(
		"POST",
		"/v1/product",
		bytes.NewBuffer(productJson),
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
	var originalProduct model.Product
	json.Unmarshal(bodyBytes, &originalProduct)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/product/%d", originalProduct.Id),
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
	var responseProduct model.Product
	json.Unmarshal(bodyBytes, &responseProduct)

	if !reflect.DeepEqual(responseProduct, originalProduct) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseProduct,
				originalProduct,
			),
		)
	}
}

func TestProductCreateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	product := model.Product{
		ProductCategoryId: 2,
		Name:              "",
		Description:       "Product Test description.",
		Price:             1.99,
		Deposit:           0.99,
		Duration:          0.5,
	}

	productJson, err := json.Marshal(product)

	req, err := http.NewRequest(
		"POST",
		"/v1/product",
		bytes.NewBuffer(productJson),
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

func TestProductUpdate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	product := model.Product{
		Id:                2,
		ProductCategoryId: 2,
		Name:              "Updated Product",
		Description:       "Updated Product description.",
		Price:             1.99,
		Deposit:           0.99,
		Duration:          0.5,
	}

	productJson, err := json.Marshal(product)

	req, err := http.NewRequest(
		"PUT",
		"/v1/product",
		bytes.NewBuffer(productJson),
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
	var originalProduct model.Product
	json.Unmarshal(bodyBytes, &originalProduct)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/product/%d", originalProduct.Id),
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
	var responseProduct model.Product
	json.Unmarshal(bodyBytes, &responseProduct)

	if !reflect.DeepEqual(responseProduct, originalProduct) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseProduct,
				originalProduct,
			),
		)
	}
}

func TestProductUpdateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	product := model.Product{
		Id:                2,
		ProductCategoryId: 2,
		Name:              "",
		Description:       "Updated Product description.",
		Price:             1.99,
		Deposit:           0.99,
		Duration:          0.5,
	}

	productJson, err := json.Marshal(product)

	req, err := http.NewRequest(
		"PUT",
		"/v1/product",
		bytes.NewBuffer(productJson),
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

func TestProductDelete(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest(
		"DELETE",
		"/v1/product/1",
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
