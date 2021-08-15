package psqlstore

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlAvailabilityStore receives a pointer to an PsqlStore.
type PsqlAvailabilityStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlAvailabilityStore.
func (s *PsqlStore) Availability() store.AvailabilityStore {
	return &PsqlAvailabilityStore{s}
}

// Get all rows in the 'availability' pg table.
//
// Returns a slice of Availability structs, and any errors encountered.
func (s *PsqlAvailabilityStore) GetAll() ([]model.Availability, error) {
	var availabilitys []model.Availability = []model.Availability{}
	rows, err := s.db.Query(`
		SELECT
			id,
			account_id,
			start_date,
			end_date
		FROM
			availability
		LIMIT 10000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'availability' rows")
		log.Println(err)
		return []model.Availability{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var availability model.Availability
		err = rows.Scan(
			&availability.Id,
			&availability.AccountId,
			&availability.StartDate,
			&availability.EndDate,
		)
		if err != nil {
			log.Println("Error: Failed to populate Availability structs")
			log.Println(err)
			return []model.Availability{}, err
		}
		availabilitys = append(availabilitys, availability)
	}

	return availabilitys, nil
}

// Get a single row from the 'availability' pg table where 'id' matches the passed
// id.
//
// Returns a Availability struct, and any errors encountered.
func (s *PsqlAvailabilityStore) Get(id int) (model.Availability, error) {
	var availability model.Availability
	rows, err := s.db.Query(`
		SELECT
			id,
			account_id,
			start_date,
			end_date
		FROM
			availability
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'availability' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Availability{}, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&availability.Id,
			&availability.AccountId,
			&availability.StartDate,
			&availability.EndDate,
		)
		if err != nil {
			log.Println("Error: Failed to populate Availability struct")
			log.Println(err)
			return model.Availability{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New(
			fmt.Sprintf("Error: Failed to find 'availability' with id '%d'", id),
		)
		log.Println(err)
		return model.Availability{}, err
	}

	return availability, nil
}

// Creates a row in the 'availability' pg table using data from the passed Availability
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlAvailabilityStore) Create(availability *model.Availability) error {
	var id int
	err := s.db.QueryRow(`
		INSERT INTO availability (
			account_id,
			start_date,
			end_date
		) VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id
		;`,
		availability.Id,
		availability.AccountId,
		availability.StartDate,
		availability.EndDate,
	).Scan(&id)
	if err != nil {
		log.Println("Error: Failed to create 'availability' row")
		log.Println(err)
		return err
	}

	availability.Id = id

	return nil
}

// Updates a row in the 'availability' pg table using data from the passed Availability
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlAvailabilityStore) Update(availability *model.Availability) error {
	_, err := s.db.Exec(`
		UPDATE
			availability
		SET
			email = $2,
			start_date = $3,
			end_date = $4
		WHERE
			id = $1
		;`,
		availability.Id,
		availability.AccountId,
		availability.StartDate,
		availability.EndDate,
	)
	if err != nil {
		log.Println("Error: Failed to update 'availability' row")
		log.Println(err)
		return err
	}

	return nil
}

// Deletes a row from the 'availability' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlAvailabilityStore) Delete(id int) error {
	_, err := s.db.Exec(`
			DELETE FROM
				availability
			WHERE
				id = $1
			;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'availability' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}

	return nil
}
