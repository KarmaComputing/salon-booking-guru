package psqlstore

import (
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlPermissionStore receives a pointer to an PsqlStore.
type PsqlPermissionStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlPermissionStore.
func (s *PsqlStore) Permission() store.PermissionStore {
	return &PsqlPermissionStore{s}
}

// Get a single row from the 'permission' pg table where 'id' matches the
// passed id.
//
// Returns Permission struct, and any errors encountered.
func (s *PsqlPermissionStore) Get(id int) (model.Permission, error) {
	var permission model.Permission
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			permission
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'permission' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Permission{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&permission.Id,
			&permission.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate Permission struct'")
			log.Println(err)
			return model.Permission{}, err
		}
	}

	return permission, nil
}

// Get all rows in the 'permission' pg table.
//
// Returns slice of Permission structs, and any errors encountered.
func (s *PsqlPermissionStore) GetAll() ([]model.Permission, error) {
	var permissions []model.Permission
	rows, err := s.db.Query(`
		SELECT
			id,
			name
		FROM
			permission
		LIMIT 100
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'permission' rows")
		log.Println(err)
		return []model.Permission{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission model.Permission
		err = rows.Scan(
			&permission.Id,
			&permission.Name,
		)
		if err != nil {
			log.Println("Error: Failed to populate Permission structs'")
			log.Println(err)
			return []model.Permission{}, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

// Creates a row in the 'permission' pg table using data from the passed Permission
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlPermissionStore) Create(permission *model.Permission) error {
	var id int
	err := s.db.QueryRow(`
		INSERT INTO permission (name)
		SELECT name FROM permission
		UNION
		VALUES ($1)
		EXCEPT
		SELECT name FROM permission
		RETURNING id
		;`,
		permission.Name,
	).Scan(&id)
	if err != nil {
		log.Println("Error: Failed to create 'permission' row")
		log.Println(err)
		return err
	}

	permission.Id = id
	return nil
}

// Updates a row in the 'permission' pg table using data from the passed Permission
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlPermissionStore) Update(permission *model.Permission) error {
	_, err := s.db.Exec(`
		UPDATE
			permission
		SET
			name=$1
		WHERE
			id=$2
		;`,
		permission.Name,
		permission.Id,
	)
	if err != nil {
		log.Println("Error: Failed to update 'permission' row")
		log.Println(err)
		return err
	}
	return nil
}

// Deletes a row from the 'permission' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlPermissionStore) Delete(id int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			role_permission_link
		WHERE
			permission_id = $1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'role_permission_link' rows with permission_id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}

	_, err = s.db.Exec(`
				DELETE FROM
					permission
				WHERE
					id = $1
				;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'permission' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}
	return nil
}

// Get all rows in the 'permission' pg table, based on a relationship with
// 'account' table's 'role_id' where 'account.id' matches the
// passed accountId.
//
// Returns slice of strings, and any errors encountered.
func (s *PsqlPermissionStore) GetAllNameByAccountId(accountId int) ([]string, error) {
	var permissions []string = []string{}
	rows, err := s.db.Query(`
		SELECT
			p.name
		FROM
			permission AS p
		INNER JOIN
			role_permission_link AS rpl
		ON
			p.id = rpl.permission_id
		INNER JOIN
			account AS a
		ON
			rpl.role_id = a.role_id
		WHERE
			a.id = $1
		GROUP BY
			p.id
		LIMIT 100
		;`,
		accountId,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'permission' rows")
		log.Println(err)
		return []string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission string
		err = rows.Scan(
			&permission,
		)
		if err != nil {
			log.Println("Error: Failed to populate permissions slice'")
			log.Println(err)
			return []string{}, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
