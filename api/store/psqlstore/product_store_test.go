package psqlstore

import (
	"errors"
	"fmt"
	"reflect"
	"salon-booking-guru/store/model"
	"testing"
)

func TestProductGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	products, err := s.Product().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 4 {
		t.Fatal(errors.New("Number of products returned is invalid"))
	}
}

func TestProductGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	product := model.Product{
		Id:                2,
		ProductCategoryId: 3,
		Name:              "Product 2",
		Description:       "Product 2 description.",
		Price:             24.99,
		Deposit:           4.50,
		Duration:          2.5,
	}

	productGet, err := s.Product().Get(product.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(product, productGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", product, productGet))
	}
}

func TestProductCreate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	product := model.Product{
		ProductCategoryId: 2,
		Name:              "Product Test",
		Description:       "Product Test description.",
		Price:             1.99,
		Deposit:           0.99,
		Duration:          0.5,
	}

	err = s.Product().Create(&product)
	if err != nil {
		t.Fatal(err)
	}

	productGet, err := s.Product().Get(product.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(product, productGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", product, productGet))
	}
}

func TestProductUpdate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	product := model.Product{
		Id:                2,
		ProductCategoryId: 2,
		Name:              "Updated Product",
		Description:       "Updated Product description.",
		Price:             1.99,
		Deposit:           0.99,
		Duration:          0.5,
	}

	err = s.Product().Update(&product)
	if err != nil {
		t.Fatal(err)
	}

	productGet, err := s.Product().Get(product.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(product, productGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", product, productGet))
	}
}

func TestProductDelete(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.Product().Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	products, err := s.Product().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 3 {
		t.Fatal(errors.New("Number of products returned is invalid"))
	}
}
