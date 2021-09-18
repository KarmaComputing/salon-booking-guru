package psqlstore

import (
	"errors"
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

/* func TestProductGetAllNameByAccountId(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	productNames, err := s.Product().GetAllNameByAccountId(2)
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := []string{
		"Product 2",
		"Product 3",
		"Product 4",
	}

	if !reflect.DeepEqual(productNames, expectedOutput) {
		t.Fatal(errors.New("Products returned are invalid"))
	}
}

func TestProductGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	product := model.Product{
		Id:   2,
		Name: "Product 2",
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
		Name: "Test Product",
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
		Id:   2,
		Name: "Updated Product Name",
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
} */
