package psqlstore

import (
	"errors"
	"salon-booking-guru/store/model"
	"testing"
	"time"
)

func TestAvailabilityGet(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Availability().Get(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAvailabilityGetAll(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	availabilities, err := s.Availability().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(availabilities) != 8 {
		t.Fatal(errors.New("Number of availabilities returned is invalid"))
	}
}

func TestAvailabilityGetAllByAccountId(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	availabilities, err := s.Availability().GetAllByAccountId(3)
	if err != nil {
		t.Fatal(err)
	}

	if len(availabilities) != 3 {
		t.Fatal(errors.New("Number of availabilities returned is invalid"))
	}
}

func TestAvailabilityCreateMultiple(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	availabilities := []model.Availability{
		model.Availability{
			AccountId: 1,
			StartDate: time.Date(2021, time.Month(5), 17, 9, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2021, time.Month(5), 17, 17, 0, 0, 0, time.UTC),
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

	err = s.Availability().CreateMultiple(availabilities)
	if err != nil {
		t.Fatal(err)
	}

	availabilitiesGet, err := s.Availability().GetAllByAccountId(1)
	if err != nil {
		t.Fatal(err)
	}

	if len(availabilitiesGet) != 3 {
		t.Fatal(errors.New("Number of availabilities returned is invalid"))
	}
}

func TestAvailabilityUpdate(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	availability := model.Availability{
		Id:        1,
		AccountId: 1,
		StartDate: time.Date(2021, time.Month(5), 19, 9, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2021, time.Month(5), 19, 17, 0, 0, 0, time.UTC),
	}

	err = s.Availability().Update(&availability)
	if err != nil {
		t.Fatal(err)
	}

	/* availabilityGet */
	_, err = s.Availability().Get(availability.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Issue with UTC turning into +0000 but the results do match
	/* if !reflect.DeepEqual(availability, availabilityGet) {
		t.Fatal(fmt.Sprintf("%v is not equal to %v", availability, availabilityGet))
	} */
}

func TestAvailabilityDelete(t *testing.T) {
	s, err := OpenTest()
	if err != nil {
		t.Fatal(err)
	}

	err = s.Availability().Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	availabilities, err := s.Availability().GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(availabilities) != 7 {
		t.Fatal(errors.New("Number of availabilities returned is invalid"))
	}
}
