package psqlstore

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"

	"github.com/lib/pq"
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
	var accounts []model.Account = []model.Account{}
	rows, err := s.db.Query(`
		SELECT
			id,
			role_id,
			email,
			first_name,
			last_name,
			mobile_number
		FROM
			account
		LIMIT 10000
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
			&account.RoleId,
			&account.Email,
			&account.FirstName,
			&account.LastName,
			&account.MobileNumber,
		)
		if err != nil {
			log.Println("Error: Failed to populate Account structs")
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
			role_id,
			mobile_number
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

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&account.Id,
			&account.Email,
			&account.FirstName,
			&account.LastName,
			&account.RoleId,
			&account.MobileNumber,
		)
		if err != nil {
			log.Println("Error: Failed to populate Account struct")
			log.Println(err)
			return model.Account{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New(
			fmt.Sprintf("Error: Failed to find 'account' with id '%d'", id),
		)
		log.Println(err)
		return model.Account{}, err
	}

	return account, nil
}

// Get a single row from the 'account' pg table where 'id' matches the passed
// id.
//
// Returns a AccountInfo struct, and any errors encountered.
func (s *PsqlAccountStore) GetInfo(id int) (model.AccountInfo, error) {
	var accountInfo model.AccountInfo
	rows, err := s.db.Query(`
		SELECT
			account.email,
			account.first_name,
			account.last_name,
			role.name,
			ARRAY(
				SELECT
					permission.name
				FROM
					role_permission_link
				INNER JOIN
					permission
				ON
					permission.id = role_permission_link.permission_id
				WHERE
					role_id = role.id
			)
		FROM
			account
		INNER JOIN
			role
		ON
			account.role_id = role.id
		WHERE
			account.id = $1
		LIMIT 1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to find 'account' with id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return model.AccountInfo{}, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		err = rows.Scan(
			&accountInfo.Email,
			&accountInfo.FirstName,
			&accountInfo.LastName,
			&accountInfo.RoleName,
			pq.Array(&accountInfo.Permissions),
		)
		if err != nil {
			log.Println("Error: Failed to populate AccountInfo struct")
			log.Println(err)
			return model.AccountInfo{}, err
		}
		counter++
	}

	if counter == 0 {
		err = errors.New(
			fmt.Sprintf("Error: Failed to find 'account' with id '%d'", id),
		)
		log.Println(err)
		return model.AccountInfo{}, err
	}

	return accountInfo, nil
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
			role_id,
			mobile_number
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
		account.Email,
		account.FirstName,
		account.LastName,
		string(hashedPass),
		account.RoleId,
		account.MobileNumber,
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
				mobile_number = $7
			WHERE
				id = $1
			;`,
			account.Id,
			account.Email,
			account.FirstName,
			account.LastName,
			hashedPass,
			account.RoleId,
			account.MobileNumber,
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
				mobile_number = $6
			WHERE
				id = $1
			;`,
			account.Id,
			account.Email,
			account.FirstName,
			account.LastName,
			account.RoleId,
			account.MobileNumber,
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

// Deletes all account_qualification_link rows where account_id matches the
// passed accountId, then inserts new links rows based on the passed
// qualificationIds.
//
// Returns any errors encountered.
func (s *PsqlAccountStore) UpsertQualification(id int, qualificationIds []int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			account_qualification_link
		WHERE
			account_id = $1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'account_qualification_link' rows with account_id '" + strconv.Itoa(id) + "'")
		log.Println(err)
		return err
	}

	sqlValues := ""
	for i, qualificationId := range qualificationIds {
		sqlValues += fmt.Sprintf("(%d, %d)", id, qualificationId)
		if i != len(qualificationIds)-1 {
			sqlValues += ","
		}
	}

	_, err = s.db.Exec(
		fmt.Sprintf(`
			INSERT INTO account_qualification_link (
				account_id,
				qualification_id
			) VALUES
				%s
			;`,
			sqlValues,
		),
	)
	if err != nil {
		log.Println("Error: Failed to insert 'account_qualification_link' rows")
		log.Println(err)
		return err
	}

	return nil
}
