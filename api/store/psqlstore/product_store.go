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

// Get all rows in the 'product' pg table in summary form.
//
// Returns a slice of ProductSummary structs, and any errors encountered.
func (s *PsqlProductStore) GetAllSummary() ([]model.ProductSummary, error) {
	var productSummaries []model.ProductSummary
	rows, err := s.db.Query(`
		SELECT
			p.id,
			pc.name,
			p.name,
			p.description,
			p.price,
			p.deposit,
			p.duration
		FROM
			product AS p
		INNER JOIN
			product_category AS pc
		ON
			p.product_category_id = pc.id
		LIMIT 10000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'productSummary' rows")
		log.Println(err)
		return []model.ProductSummary{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var productSummary model.ProductSummary
		err = rows.Scan(
			&productSummary.Id,
			&productSummary.ProductCategoryName,
			&productSummary.Name,
			&productSummary.Description,
			&productSummary.Price,
			&productSummary.Deposit,
			&productSummary.Duration,
		)
		if err != nil {
			log.Println("Error: Failed to populate ProductSummary structs")
			log.Println(err)
			return []model.ProductSummary{}, err
		}
		productSummaries = append(productSummaries, productSummary)
	}

	return productSummaries, nil
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
			product_category_id,
			name,
			description,
			price,
			deposit,
			duration
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
			&product.ProductCategoryId,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Deposit,
			&product.Duration,
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
			product_category_id,
			name,
			description,
			price,
			deposit,
			duration
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
		RETURNING id
		;`,
		product.ProductCategoryId,
		product.Name,
		product.Description,
		product.Price,
		product.Deposit,
		product.Duration,
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
			product_category_id = $2,
			name = $3,
			description = $4,
			price = $5,
			deposit = $6,
			duration = $7
		WHERE
			id = $1
		;`,
		product.Id,
		product.ProductCategoryId,
		product.Name,
		product.Description,
		product.Price,
		product.Deposit,
		product.Duration,
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

// Deletes all product_qualification_link rows where product_id matches the
// passed productId, then inserts new links rows based on the passed
// qualificationIds.
//
// Returns any errors encountered.
func (s *PsqlProductStore) UpsertQualification(productId int, qualificationIds []int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			product_qualification_link
		WHERE
			product_id = $1
		;`,
		productId,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'product_qualification_link' rows with product_id '" + strconv.Itoa(productId) + "'")
		log.Println(err)
		return err
	}

	sqlValues := ""
	for i, qualificationId := range qualificationIds {
		sqlValues += fmt.Sprintf("(%d, %d)", productId, qualificationId)
		if i != len(qualificationIds)-1 {
			sqlValues += ","
		}
	}

	_, err = s.db.Exec(
		fmt.Sprintf(`
			INSERT INTO product_qualification_link (
				product_id,
				qualification_id
			) VALUES
				%s
			;`,
			sqlValues,
		),
	)
	if err != nil {
		log.Println("Error: Failed to insert 'product_qualification_link' rows")
		log.Println(err)
		return err
	}

	return nil
}
