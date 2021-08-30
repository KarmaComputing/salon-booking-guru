package psqlstore

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlProductCategoryStore receives a pointer to an PsqlStore.
type PsqlProductCategoryStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlProductCategoryStore.
func (s *PsqlStore) ProductCategory() store.ProductCategoryStore {
	return &PsqlProductCategoryStore{s}
}

// Get all rows in the 'productCategory' pg table.
//
// Returns a slice of ProductCategory structs, and any errors encountered.
func (s *PsqlProductCategoryStore) GetAll() ([]model.ProductCategory, error) {
	var productCategories []model.ProductCategory
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			product_category
		LIMIT 10000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'product_category' rows")
		log.Println(err)
		return []model.ProductCategory{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var productCategory model.ProductCategory
		err = rows.Scan(
			&productCategory.Id,
			&productCategory.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate ProductCategory structs")
			log.Println(err)
			return []model.ProductCategory{}, err
		}
		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

// Get a single row from the 'productCategory' pg table where 'id' matches the passed
// id.
//
// Returns a ProductCategory struct, and any errors encountered.
func (s *PsqlProductCategoryStore) Get(id int) (model.ProductCategory, error) {
	var productCategory model.ProductCategory
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			product_category
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'product_category' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.ProductCategory{}, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&productCategory.Id,
			&productCategory.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate ProductCategory struct'")
			log.Println(err)
			return model.ProductCategory{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New(
			fmt.Sprintf("Error: Failed to find 'product_category' with id '%d'", id),
		)
		log.Println(err)
		return model.ProductCategory{}, err
	}

	return productCategory, nil
}

// Creates a row in the 'productCategory' pg table using data from the passed ProductCategory
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlProductCategoryStore) Create(productCategory *model.ProductCategory) error {
	var id int
	err := s.db.QueryRow(`
		INSERT INTO product_category (
			name
		) VALUES (
			$1
		)
		RETURNING id
		;`,
		productCategory.Name,
	).Scan(&id)
	if err != nil {
		log.Println("Error: Failed to create 'product_category' row")
		log.Println(err)
		return err
	}

	productCategory.Id = id

	return nil
}

// Updates a row in the 'productCategory' pg table using data from the passed ProductCategory
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlProductCategoryStore) Update(productCategory *model.ProductCategory) error {
	_, err := s.db.Exec(`
		UPDATE
			product_category
		SET
			name = $2
		WHERE
			id = $1
		;`,
		productCategory.Id,
		productCategory.Name,
	)
	if err != nil {
		log.Println("Error: Failed to update 'product_category' row")
		log.Println(err)
		return err
	}

	return nil
}

// Deletes a row from the 'productCategory' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlProductCategoryStore) Delete(id int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			product_category
		WHERE
			id = $1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'product_category' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}
	return nil
}
