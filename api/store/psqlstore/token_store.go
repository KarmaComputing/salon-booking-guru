package psqlstore

import (
	"log"
	"strconv"

	"salon-booking-guru/crypto"
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

// Generates a token and inserts it into the database.
//
// Returns the token and any errors encountered.
func (s *PsqlTokenStore) Generate(accountId int) (model.Token, error) {
	token, err := crypto.GenerateToken()
	if err != nil {
		return model.Token{}, err
	}

	_, err = s.db.Exec(`
		INSERT INTO token (
			token,
			account_id
		) VALUES (
			$1,
			$2
		)
		;`,
		token,
		accountId,
	)
	if err != nil {
		log.Println("Error: Failed to insert 'token' row")
		log.Println(err)
		return model.Token{}, err
	}

	return model.Token{
		Token:     token,
		AccountId: accountId,
	}, nil
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
