package validation

import (
	"errors"
	"salon-booking-guru/store/model"
)

func ValidateProductCategory(productCategory model.ProductCategory) error {
	var v Validation

	v.minLength("Name", productCategory.Name, 1)
	v.maxLength("Name", productCategory.Name, 200)

	if !v.IsValid() {
		return errors.New("Error: ProductCategory data is not valid")
	}

	return nil
}
