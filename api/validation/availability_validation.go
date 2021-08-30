package validation

import (
	"errors"
	"salon-booking-guru/store/model"
)

func ValidateAvailability(availability model.Availability) error {
	var v Validation

	v.isSelected("Account", availability.AccountId)

	if !v.IsValid() {
		return errors.New("Error: Availability data is not valid")
	}

	return nil
}
