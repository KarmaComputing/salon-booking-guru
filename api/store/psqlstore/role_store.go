package psqlstore

import (
	"log"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlRoleStore receives a pointer to an PsqlStore.
type PsqlRoleStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlRoleStore.
func (s *PsqlStore) Role() store.RoleStore {
	return &PsqlRoleStore{s}
}

// Get all rows in the 'role' pg table.
//
// Returns a slice of Role structs, and any errors encountered.
func (s *PsqlRoleStore) GetAll() ([]model.Role, error) {
	var roles []model.Role = []model.Role{}
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			role
		LIMIT 10000
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
		)
		if err != nil {
			log.Println("Error: Failed to populate Role structs")
			log.Println(err)
			return []model.Role{}, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
