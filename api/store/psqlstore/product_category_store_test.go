package psqlstore

import (
	"errors"
	"fmt"
	"reflect"
	"salon-booking-guru/store/model"
	"testing"
)

func TestProductCategoryGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	productCategories, err := s.ProductCategory().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(productCategories) != 3 {
		t.Fatal(errors.New("Number of productCategories returned is invalid"))
	}
}

func TestProductCategoryGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	productCategory := model.ProductCategory{
		Id:   2,
		Name: "Product Category 2",
	}

	productCategoryGet, err := s.ProductCategory().Get(productCategory.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(productCategory, productCategoryGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", productCategory, productCategoryGet))
	}
}

func TestProductCategoryCreate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	productCategory := model.ProductCategory{
		Name: "Test ProductCategory",
	}

	err = s.ProductCategory().Create(&productCategory)
	if err != nil {
		t.Fatal(err)
	}

	productCategoryGet, err := s.ProductCategory().Get(productCategory.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(productCategory, productCategoryGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", productCategory, productCategoryGet))
	}
}

func TestProductCategoryUpdate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	productCategory := model.ProductCategory{
		Id:   2,
		Name: "Updated ProductCategory Name",
	}

	err = s.ProductCategory().Update(&productCategory)
	if err != nil {
		t.Fatal(err)
	}

	productCategoryGet, err := s.ProductCategory().Get(productCategory.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(productCategory, productCategoryGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", productCategory, productCategoryGet))
	}
}

func TestProductCategoryDelete(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.ProductCategory().Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	productCategories, err := s.ProductCategory().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(productCategories) != 2 {
		t.Fatal(errors.New("Number of productCategories returned is invalid"))
	}
}
