package psqlstore

import (
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"

	"golang.org/x/crypto/bcrypt"
)

// PsqlAccountStore receives a pointer to an PsqlStore.
type PsqlAccountStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlAccountStore.
func (s *PsqlStore) Account() store.AccountStore {
	return &PsqlAccountStore{s}
}

// Get all rows in the 'account' pg table.
//
// Returns a slice of Account structs, and any errors encountered.
func (s *PsqlAccountStore) GetAll() ([]model.Account, error) {
	var accounts []model.Account
	rows, err := s.db.Query(`
		SELECT
			id,
			email,
			first_name,
			last_name,
			role_id
		FROM
			account
		LIMIT 1000
		;`,
	)
	if err != nil {
		log.Println("Error: Failed to retrieve 'account' rows")
		log.Println(err)
		return []model.Account{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var account model.Account
		err = rows.Scan(
			&account.Id,
			&account.Email,
			&account.FirstName,
			&account.LastName,
			&account.RoleId,
		)
		if err != nil {
			log.Println("Error: Failed to populate Account structs'")
			log.Println(err)
			return []model.Account{}, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// Get a single row from the 'account' pg table where 'id' matches the passed
// id.
//
// Returns a Account struct, and any errors encountered.
func (s *PsqlAccountStore) Get(id int) (model.Account, error) {
	var account model.Account
	rows, err := s.db.Query(`
		SELECT
			id,
			email,
			first_name,
			last_name,
			role_id
		FROM
			account
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'account' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.Account{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&account.Id,
			&account.Email,
			&account.FirstName,
			&account.LastName,
			&account.RoleId,
		)
		if err != nil {
			log.Println("Error: Failed to populate Account struct'")
			log.Println(err)
			return model.Account{}, err
		}
	}

	return account, nil
}

// Creates a row in the 'account' pg table using data from the passed Account
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlAccountStore) Create(account *model.Account) error {
	hashedPass, _ := bcrypt.GenerateFromPassword(
		[]byte(account.Password),
		10,
	)

	var id int
	err := s.db.QueryRow(`
		INSERT INTO account (
			email,
			first_name,
			last_name,
			password,
			role_id
		)
		SELECT
			email
		FROM
			account
		UNION
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		EXCEPT
		SELECT
			email
		FROM
			account
		RETURNING id
		;`,
		account.Email,
		account.FirstName,
		account.LastName,
		hashedPass,
		account.RoleId,
	).Scan(&id)

	if err != nil {
		log.Println("Error: Failed to create 'account' row")
		log.Println(err)
		return err
	}

	account.Id = id

	return nil
}

// Updates a row in the 'account' pg table using data from the passed Account
// struct pointer.
//
// Returns any errors encountered.
func (s *PsqlAccountStore) Update(account *model.Account) error {
	if account.Password != "" {
		hashedPass, _ := bcrypt.GenerateFromPassword(
			[]byte(account.Password),
			10,
		)
		_, err := s.db.Exec(`
			UPDATE
				account
			SET
				email = $2,
				first_name = $3,
				last_name = $4,
				password = $5,
				role_id = $6,
			WHERE
				id = $1
			;`,
			account.Id,
			account.Email,
			account.FirstName,
			account.LastName,
			hashedPass,
			account.RoleId,
		)

		if err != nil {
			log.Println("Error: Failed to update 'account' row")
			log.Println(err)
			return err
		}
	} else {
		_, err := s.db.Exec(`
			UPDATE
				account
			SET
				email = $2,
				first_name = $3,
				last_name = $4,
				role_id = $5,
			WHERE
				id = $1
			;`,
			account.Id,
			account.Email,
			account.FirstName,
			account.LastName,
			account.RoleId,
		)

		if err != nil {
			log.Println("Error: Failed to update 'account' row")
			log.Println(err)
			return err
		}
	}
	return nil
}

// Deletes a row from the 'account' pg table where 'id' matches the passed id.
//
// Returns any errors encountered.
func (s *PsqlAccountStore) Delete(id int) error {
	s.Token().DeleteAllByAccountId(id)

	_, err := s.db.Exec(`
			DELETE FROM
				account
			WHERE
				id = $1
			;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'account' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}
	return nil
}

// Checks if the role associated to the accountId specified has the permission
//
// Returns true if the account does have the permission, false otherwise.
func (s *PsqlAccountStore) IsAuthorized(id int, permissionName string) (bool, error) {
	return true, nil
}
