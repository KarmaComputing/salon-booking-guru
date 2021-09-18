package validation

import (
	"errors"
	"salon-booking-guru/store/model"
)

func ValidateProduct(product model.Product) error {
	var v Validation

	v.isSelected("Product Category", product.ProductCategoryId)
	v.minLength("Name", product.Name, 1)
	v.maxLength("Name", product.Name, 200)
	v.minLength("Description", product.Name, 1)
	v.maxLength("Description", product.Name, 2000)

	if !v.IsValid() {
		return errors.New("Error: Product data is not valid")
	}

	return nil
}
