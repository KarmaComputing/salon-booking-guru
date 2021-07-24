package psqlstore

import (
	"errors"
	"fmt"
	"reflect"
	"salon-booking-guru/store/model"
	"testing"
)

func TestAccountGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	accounts, err := s.Account().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) != 4 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}

func TestAccountGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	account := model.Account{
		Id:        2,
		RoleId:    2,
		Email:     "owner@example.com",
		FirstName: "Beatrice",
		LastName:  "Brown",
	}

	accountGet, err := s.Account().Get(account.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(account, accountGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", account, accountGet))
	}
}

func TestAccountCreate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	mobileNumber := "07123456789"
	account := model.Account{
		Email:        "test@example.com",
		FirstName:    "Test",
		LastName:     "User",
		Password:     "password",
		RoleId:       2,
		MobileNumber: &mobileNumber,
	}

	err = s.Account().Create(&account)
	if err != nil {
		t.Fatal(err)
	}

	accountGet, err := s.Account().Get(account.Id)
	if err != nil {
		t.Fatal(err)
	}

	account.Password = ""
	accountGet.Password = ""

	if !reflect.DeepEqual(account, accountGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", account, accountGet))
	}
}

func TestAccountUpdate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	mobileNumber := "07123456789"
	account := model.Account{
		Email:        "test@example.com",
		FirstName:    "Test",
		LastName:     "User",
		Password:     "password",
		RoleId:       2,
		MobileNumber: &mobileNumber,
	}

	err = s.Account().Create(&account)
	if err != nil {
		t.Fatal(err)
	}

	accountGet, err := s.Account().Get(account.Id)
	if err != nil {
		t.Fatal(err)
	}

	account.Password = ""
	accountGet.Password = ""

	if !reflect.DeepEqual(account, accountGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", account, accountGet))
	}
}

func TestAccountDelete(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.Account().Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	accounts, err := s.Account().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) != 3 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}
