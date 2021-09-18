package psqlstore

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlProductStore receives a pointer to an PsqlStore.
type PsqlProductStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlProductStore.
func (s *PsqlStore) Product() store.ProductStore {
	return &PsqlProductStore{s}
}

// Get all rows in the 'product' pg table.
//
// Returns a slice of Product structs, and any errors encountered.
func (s *PsqlProductStore) GetAll() ([]model.Product, error) {
	var products []model.Product
	rows, err := s.db.Query(`
		SELECT
			id,
			product_category_id,
			name,
			description,
			price,
			deposit,
			duration
		FROM
			product
		LIMIT 10000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'product' rows")
		log.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.Id,
			&product.ProductCategoryId,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Deposit,
			&product.Duration,
		)
		if err != nil {
			log.Println("Error: Failed to populate Product structs")
			log.Println(err)
			return []model.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Get all name fields in the 'product' pg table by account_id.
//
// Returns a slice of Product structs, and any errors encountered.
func (s *PsqlProductStore) GetAllNameByAccountId(accountId int) ([]string, error) {
	rows, err := s.db.Query(`
		SELECT
			name
		FROM
			product AS q
		INNER JOIN
			account_product_link AS aql
		ON
			aql.product_id = q.id
		WHERE
			aql.account_id = $1
		LIMIT 10000
		;`,
		accountId,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'product' names")
		log.Println(err)
		return []string{}, err
	}
	defer rows.Close()

	var productNames []string = []string{}
	for rows.Next() {
		var productName string
		err = rows.Scan(
			&productName,
		)
		if err != nil {
			log.Println("Error: Failed to populate Product structs")
			log.Println(err)
			return []string{}, err
		}
		productNames = append(productNames, productName)
	}

	return productNames, nil
}

// Get a single row from the 'product' pg table where 'id' matches the passed
// id.
//
// Returns a Product struct, and any errors encountered.
func (s *PsqlProductStore) Get(id int) (model.Product, error) {
	var product model.Product
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			product
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'product' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Product{}, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate Product struct'")
			log.Println(err)
			return model.Product{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New(
			fmt.Sprintf("Error: Failed to find 'product' with id '%d'", id),
		)
		log.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

// Creates a row in the 'product' pg table using data from the passed Product
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlProductStore) Create(product *model.Product) error {
	var id int
	err := s.db.QueryRow(`
		INSERT INTO product (
			name
		) VALUES (
			$1
		)
		RETURNING id
		;`,
		product.Name,
	).Scan(&id)
	if err != nil {
		log.Println("Error: Failed to create 'product' row")
		log.Println(err)
		return err
	}

	product.Id = id

	return nil
}

// Updates a row in the 'product' pg table using data from the passed Product
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlProductStore) Update(product *model.Product) error {
	_, err := s.db.Exec(`
		UPDATE
			product
		SET
			name = $2
		WHERE
			id = $1
		;`,
		product.Id,
		product.Name,
	)
	if err != nil {
		log.Println("Error: Failed to update 'product' row")
		log.Println(err)
		return err
	}

	return nil
}

// Deletes a row from the 'product' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlProductStore) Delete(id int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			product
		WHERE
			id = $1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'product' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}
	return nil
}
