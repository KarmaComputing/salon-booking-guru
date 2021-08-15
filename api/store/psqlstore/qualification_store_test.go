package psqlstore

import (
	"errors"
	"fmt"
	"reflect"
	"salon-booking-guru/store/model"
	"testing"
)

func TestQualificationGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	qualifications, err := s.Qualification().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(qualifications) != 4 {
		t.Fatal(errors.New("Number of qualifications returned is invalid"))
	}
}

func TestQualificationGetAllNameByAccountId(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	qualificationNames, err := s.Qualification().GetAllNameByAccountId(2)
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := []string{
		"Qualification 2",
		"Qualification 3",
		"Qualification 4",
	}

	if !reflect.DeepEqual(qualificationNames, expectedOutput) {
		t.Fatal(errors.New("Qualifications returned are invalid"))
	}
}

func TestQualificationGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	qualification := model.Qualification{
		Id:   2,
		Name: "Qualification 2",
	}

	qualificationGet, err := s.Qualification().Get(qualification.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(qualification, qualificationGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", qualification, qualificationGet))
	}
}

func TestQualificationCreate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	qualification := model.Qualification{
		Name: "Test Qualification",
	}

	err = s.Qualification().Create(&qualification)
	if err != nil {
		t.Fatal(err)
	}

	qualificationGet, err := s.Qualification().Get(qualification.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(qualification, qualificationGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", qualification, qualificationGet))
	}
}

func TestQualificationUpdate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	qualification := model.Qualification{
		Id:   2,
		Name: "Updated Qualification Name",
	}

	err = s.Qualification().Update(&qualification)
	if err != nil {
		t.Fatal(err)
	}

	qualificationGet, err := s.Qualification().Get(qualification.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(qualification, qualificationGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", qualification, qualificationGet))
	}
}

func TestQualificationDelete(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.Qualification().Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	qualifications, err := s.Qualification().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(qualifications) != 3 {
		t.Fatal(errors.New("Number of qualifications returned is invalid"))
	}
}
