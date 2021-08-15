package psqlstore

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlQualificationStore receives a pointer to an PsqlStore.
type PsqlQualificationStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlQualificationStore.
func (s *PsqlStore) Qualification() store.QualificationStore {
	return &PsqlQualificationStore{s}
}

// Get all rows in the 'qualification' pg table.
//
// Returns a slice of Qualification structs, and any errors encountered.
func (s *PsqlQualificationStore) GetAll() ([]model.Qualification, error) {
	var qualifications []model.Qualification
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			qualification
		LIMIT 10000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'qualification' rows")
		log.Println(err)
		return []model.Qualification{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var qualification model.Qualification
		err = rows.Scan(
			&qualification.Id,
			&qualification.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate Qualification structs")
			log.Println(err)
			return []model.Qualification{}, err
		}
		qualifications = append(qualifications, qualification)
	}

	return qualifications, nil
}

// Get a single row from the 'qualification' pg table where 'id' matches the passed
// id.
//
// Returns a Qualification struct, and any errors encountered.
func (s *PsqlQualificationStore) Get(id int) (model.Qualification, error) {
	var qualification model.Qualification
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			qualification
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'qualification' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Qualification{}, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&qualification.Id,
			&qualification.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate Qualification struct'")
			log.Println(err)
			return model.Qualification{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New(
			fmt.Sprintf("Error: Failed to find 'qualification' with id '%d'", id),
		)
		log.Println(err)
		return model.Qualification{}, err
	}

	return qualification, nil
}

// Creates a row in the 'qualification' pg table using data from the passed Qualification
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlQualificationStore) Create(qualification *model.Qualification) error {
	var id int
	err := s.db.QueryRow(`
		INSERT INTO qualification (
			name
		) VALUES (
			$1
		)
		RETURNING id
		;`,
		qualification.Name,
	).Scan(&id)
	if err != nil {
		log.Println("Error: Failed to create 'qualification' row")
		log.Println(err)
		return err
	}

	qualification.Id = id

	return nil
}

// Updates a row in the 'qualification' pg table using data from the passed Qualification
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlQualificationStore) Update(qualification *model.Qualification) error {
	_, err := s.db.Exec(`
		UPDATE
			qualification
		SET
			name = $2
		WHERE
			id = $1
		;`,
		qualification.Id,
		qualification.Name,
	)
	if err != nil {
		log.Println("Error: Failed to update 'qualification' row")
		log.Println(err)
		return err
	}

	return nil
}

// Deletes a row from the 'qualification' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlQualificationStore) Delete(id int) error {
	_, err := s.db.Exec(`
			DELETE FROM
				qualification
			WHERE
				id = $1
			;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'qualification' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}
	return nil
}
