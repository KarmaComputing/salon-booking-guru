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

	if len(accounts) != 5 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}

func TestAccountGetAllSummary(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	accountSummaries, err := s.Account().GetAllSummary()
	if err != nil {
		t.Fatal(err)
	}

	if len(accountSummaries) != 5 {
		t.Fatal(errors.New("Number of account summaries returned is invalid"))
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

func TestAccountGetInfo(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	accountInfo := model.AccountInfo{
		Email:       "permissiontest@example.com",
		FirstName:   "Edgar",
		LastName:    "Evans",
		RoleName:    "PermissionTest",
		Permissions: []string{"canPermissionTest"},
	}

	accountInfoGet, err := s.Account().GetInfo(5)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(accountInfo, accountInfoGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", accountInfo, accountInfoGet))
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
		Id:           2,
		Email:        "test@example.com",
		FirstName:    "Test",
		LastName:     "User",
		Password:     "password",
		RoleId:       2,
		MobileNumber: &mobileNumber,
	}

	err = s.Account().Update(&account)
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

func TestAccountUpsertQualification(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.Account().UpsertQualification(2, []int{1, 2})
	if err != nil {
		t.Fatal(err)
	}

	qualificationNames, err := s.Qualification().GetAllNameByAccountId(2)
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := []string{
		"Qualification 1",
		"Qualification 2",
	}

	if !reflect.DeepEqual(qualificationNames, expectedOutput) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", qualificationNames, expectedOutput))

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

	if len(accounts) != 4 {
		t.Fatal(errors.New("Number of accounts returned is invalid"))
	}
}
