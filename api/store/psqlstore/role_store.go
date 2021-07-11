package psqlstore

import (
	"errors"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"

	"github.com/lib/pq"
)

// PsqlRoleStore receives a pointer to an PsqlStore.
type PsqlRoleStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlRoleStore.
func (s *PsqlStore) Role() store.RoleStore {
	return &PsqlRoleStore{s}
}

// Get a single row from the 'role' pg table where 'id' matches the passed id.
//
// Returns Role struct, and any errors encountered.
func (s *PsqlRoleStore) Get(id int) (model.Role, error) {
	var role model.Role
	rows, err := s.db.Query(`
		SELECT
			id,
			name,
			ARRAY(
				SELECT
					p.name
				FROM
					permission AS p
				INNER JOIN
					role_permission_link AS rpl
				ON
					rpl.role_id = role.id
				WHERE
					rpl.permission_id = p.id
				GROUP BY
					p.name
			)
		FROM
			role
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'role' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Role{}, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&role.Id,
			&role.Name,
			pq.Array(&role.Permissions),
		)
		if err != nil {
			log.Println("Error: Failed to populate Role struct'")
			log.Println(err)
			return model.Role{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New("pq: no rows in result set")
		log.Println("Error: Failed to find 'role' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Role{}, err
	}

	return role, nil
}

// Get all rows in the 'role' pg table.
//
// Returns slice of Role structs, and any errors encountered.
func (s *PsqlRoleStore) GetAll() ([]model.Role, error) {
	var roles []model.Role
	rows, err := s.db.Query(`
		SELECT
			id,
			name,
			ARRAY(
				SELECT
					p.name
				FROM
					permission AS p
				INNER JOIN
					role_permission_link AS rpl
				ON
					rpl.role_id = role.id
				WHERE
					rpl.permission_id = p.id
				GROUP BY
					p.name
			)
		FROM
			role
		LIMIT 1000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'role' rows")
		log.Println(err)
		return []model.Role{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var role model.Role
		err = rows.Scan(
			&role.Id,
			&role.Name,
			pq.Array(&role.Permissions),
		)
		if err != nil {
			log.Println("Error: Failed to populate Role structs'")
			log.Println(err)
			return []model.Role{}, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// Creates a row in the 'role' pg table using data from the passed Role
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlRoleStore) Create(role *model.Role) error {
	var id int
	err := s.db.QueryRow(`
		INSERT INTO	role (
			name
		) VALUES (
			$1
		)
		RETURNING id
		;`,
		role.Name,
	).Scan(&id)

	if err != nil {
		log.Println("Error: Failed to create 'role' row")
		log.Println(err)
		return err
	}

	role.Id = id

	return nil
}

// Updates a row in the 'role' pg table using data from the passed Role
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlRoleStore) Update(role *model.Role) error {
	_, err := s.db.Exec(`
		UPDATE
			role
		SET
			name = $2
		WHERE
			id = $1
		`,
		role.Id,
		role.Name,
	)
	if err != nil {
		log.Println("Error: Failed to update 'role' row")
		log.Println(err)
		return err
	}

	return nil
}

// Deletes a row from the 'role' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlRoleStore) Delete(id int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			role
		WHERE
			id = $1
		;`,
		id,
	)

	if err != nil {
		log.Println("Error: Failed to delete 'role' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}

	return nil
}
