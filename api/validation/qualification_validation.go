package validation

import (
	"errors"
	"salon-booking-guru/store/model"
)

func ValidateQualification(qualification model.Qualification) error {
	var v Validation

	v.minLength("Name", qualification.Name, 1)
	v.maxLength("Name", qualification.Name, 200)

	if !v.IsValid() {
		return errors.New("Error: Qualification data is not valid")
	}

	return nil
}
