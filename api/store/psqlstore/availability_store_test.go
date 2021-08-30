package psqlstore

import (
	"errors"
	"fmt"
	"reflect"
	"salon-booking-guru/store/model"
	"testing"
)

// incomplete
func TestAvailabilityGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	accounts, err := s.Availability().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) != 5 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}

// incomplete
func TestAvailabilityGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	account := model.Availability{
		Id:        2,
		RoleId:    2,
		Email:     "owner@example.com",
		FirstName: "Beatrice",
		LastName:  "Brown",
	}

	accountGet, err := s.Availability().Get(account.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(account, accountGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", account, accountGet))
	}
}

// incomplete
func TestAvailabilityGetInfo(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	accountInfo := model.AvailabilityInfo{
		Email:       "permissiontest@example.com",
		FirstName:   "Edgar",
		LastName:    "Evans",
		RoleName:    "PermissionTest",
		Permissions: []string{"canPermissionTest"},
	}

	accountInfoGet, err := s.Availability().GetInfo(5)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(accountInfo, accountInfoGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", accountInfo, accountInfoGet))
	}
}

// incomplete
func TestAvailabilityCreate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	mobileNumber := "07123456789"
	account := model.Availability{
		Email:        "test@example.com",
		FirstName:    "Test",
		LastName:     "User",
		Password:     "password",
		RoleId:       2,
		MobileNumber: &mobileNumber,
	}

	err = s.Availability().Create(&account)
	if err != nil {
		t.Fatal(err)
	}

	accountGet, err := s.Availability().Get(account.Id)
	if err != nil {
		t.Fatal(err)
	}

	account.Password = ""
	accountGet.Password = ""

	if !reflect.DeepEqual(account, accountGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", account, accountGet))
	}
}

// incomplete
func TestAvailabilityUpdate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	mobileNumber := "07123456789"
	account := model.Availability{
		Email:        "test@example.com",
		FirstName:    "Test",
		LastName:     "User",
		Password:     "password",
		RoleId:       2,
		MobileNumber: &mobileNumber,
	}

	err = s.Availability().Create(&account)
	if err != nil {
		t.Fatal(err)
	}

	accountGet, err := s.Availability().Get(account.Id)
	if err != nil {
		t.Fatal(err)
	}

	account.Password = ""
	accountGet.Password = ""

	if !reflect.DeepEqual(account, accountGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", account, accountGet))
	}
}

// incomplete
func TestAvailabilityDelete(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.Availability().Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	accounts, err := s.Availability().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) != 4 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}
