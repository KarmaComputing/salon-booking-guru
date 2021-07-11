package psqlstore

import (
	"errors"
	"log"
	"strconv"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"
)

// PsqlTokenStore receives a pointer to an PsqlStore.
type PsqlTokenStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlTokenStore.
func (s *PsqlStore) Token() store.TokenStore {
	return &PsqlTokenStore{s}
}

// Deletes a row from the 'token' pg table where there is a match in the passed
// id.
//
// Returns any errors encountered.
func (s *PsqlTokenStore) Delete(id int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			token
		WHERE
			id = $1
		;`,
		id,
	)
	if err != nil {
		log.Println("Error: Failed to delete 'token' row")
		log.Println(err)
		return err
	}

	return nil
}

// Deletes multiple rows from the 'token' pg table where 'account_id' matches
// the passed id.
//
// Returns any errors encountered.
func (s *PsqlTokenStore) DeleteAllByAccountId(accountId int) error {
	_, err := s.db.Exec(`
		DELETE FROM
			token
		WHERE
			account_id = $1
		;`,
		accountId,
	)

	if err != nil {
		log.Println("Error: Failed to delete 'token' rows with account_id '" + strconv.Itoa(accountId) + "'")
		log.Println(err)
		return err
	}

	return nil
}

// Check the passed token exists in the database.
//
// Returns any errors encountered.
func (s *PsqlTokenStore) Check(token *model.Token) error {
	row, err := s.db.Query(`
		SELECT
			null
		FROM
			token
		WHERE
			token = $1
		AND
			account_id = $2
		;`,
		token.Token,
		token.AccountId,
	)
	defer row.Close()

	if err != nil {
		log.Println("Error: Failed to retrieve 'token' row")
		log.Println(err)
		return err
	}

	count := 0
	for row.Next() {
		count++
	}
	if count <= 0 {
		return errors.New("Token not found")
	}

	return nil
}
