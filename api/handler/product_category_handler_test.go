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

func TestProductCategoryGetAll(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/product-category", nil)
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

func TestProductCategoryGet(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest("GET", "/v1/product-category/1", nil)
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

func TestProductCategoryCreate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	productCategory := model.ProductCategory{
		Name: "Test Product Category",
	}

	productCategoryJson, err := json.Marshal(productCategory)

	req, err := http.NewRequest(
		"POST",
		"/v1/product-category",
		bytes.NewBuffer(productCategoryJson),
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
	var originalProductCategory model.ProductCategory
	json.Unmarshal(bodyBytes, &originalProductCategory)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/product-category/%d", originalProductCategory.Id),
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
	var responseProductCategory model.ProductCategory
	json.Unmarshal(bodyBytes, &responseProductCategory)

	if !reflect.DeepEqual(responseProductCategory, originalProductCategory) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseProductCategory,
				originalProductCategory,
			),
		)
	}
}

func TestProductCategoryCreateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	productCategory := model.ProductCategory{
		Name: "",
	}

	productCategoryJson, err := json.Marshal(productCategory)

	req, err := http.NewRequest(
		"POST",
		"/v1/product-category",
		bytes.NewBuffer(productCategoryJson),
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

func TestProductCategoryUpdate(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	productCategory := model.ProductCategory{
		Id:   2,
		Name: "Updated Product Category",
	}

	productCategoryJson, err := json.Marshal(productCategory)

	req, err := http.NewRequest(
		"PUT",
		"/v1/product-category",
		bytes.NewBuffer(productCategoryJson),
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
	var originalProductCategory model.ProductCategory
	json.Unmarshal(bodyBytes, &originalProductCategory)

	req, err = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/product-category/%d", originalProductCategory.Id),
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
	var responseProductCategory model.ProductCategory
	json.Unmarshal(bodyBytes, &responseProductCategory)

	if !reflect.DeepEqual(responseProductCategory, originalProductCategory) {
		t.Fatal(
			fmt.Sprintf(
				"%v is not equal to %v",
				responseProductCategory,
				originalProductCategory,
			),
		)
	}
}

func TestProductCategoryUpdateInvalidName(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	productCategory := model.ProductCategory{
		Id:   2,
		Name: "",
	}

	productCategoryJson, err := json.Marshal(productCategory)

	req, err := http.NewRequest(
		"PUT",
		"/v1/product-category",
		bytes.NewBuffer(productCategoryJson),
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

func TestProductCategoryDelete(t *testing.T) {
	s, err := psqlstore.OpenTest()
	router := mux.NewRouter()
	InitRouter(router, s)

	req, err := http.NewRequest(
		"DELETE",
		"/v1/product-category/1",
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
