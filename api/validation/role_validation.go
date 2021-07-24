package validation

import "salon-booking-guru/store/model"

func ValidateRole(role model.Role) Validation {
	var v Validation

	v.minLength("Name", role.Name, 1)
	v.maxLength("Name", role.Name, 30)

	return v
}
