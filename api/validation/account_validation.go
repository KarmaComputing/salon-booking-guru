package validation

import (
	"errors"
	"salon-booking-guru/store/model"
)

func ValidateAccount(account model.Account) error {
	var v Validation

	v.IsEmail("Email", account.Email)
	v.maxLength("Email", account.Email, 254)
	if len(account.Password) != 0 {
		v.IsValidPassword(account.Password)
	}
	v.minLength("First name", account.FirstName, 1)
	v.maxLength("First name", account.FirstName, 60)
	v.minLength("Last name", account.LastName, 1)
	v.maxLength("Last name", account.LastName, 60)

	if !v.IsValid() {
		return errors.New("Error: Account data is not valid")
	}

	return nil
}
