package psqlstore

import (
	"log"
	"strings"

	"salon-booking-guru/crypto"
	"salon-booking-guru/store"
	"salon-booking-guru/store/model"

	"golang.org/x/crypto/bcrypt"
)

// PsqlAuthStore receives a pointer to an PsqlStore.
type PsqlAuthStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlAuthStore.
func (s *PsqlStore) Auth() store.AuthStore {
	return &PsqlAuthStore{s}
}

// Checks if the passed account has credentials that match a row in the
// database and if so, generates a token, and adds it to the database.
//
// Returns the generated token, and any errors encountered.
func (s *PsqlAuthStore) LogIn(account *model.Account) (model.Token, error) {
	var hashedPass string
	err := s.db.QueryRow(`
		SELECT
			id,
			password
		FROM
			account
		WHERE
			email=$1
		;`,
		strings.ToLower(account.Email),
	).Scan(
		&account.Id,
		&hashedPass,
	)
	if err != nil {
		log.Println("Error: Failed to find 'account' with matching id and hashed password")
		log.Println(err)
		return model.Token{}, err
	}

	var tokenModel model.Token
	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPass),
		[]byte(account.Password),
	)
	if err != nil {
		log.Println("Error: Failed login attempt by '" + account.Email + "'")
		log.Println(err)
		return model.Token{}, err
	}

	token, err := crypto.GenerateToken()
	if err != nil {
		log.Println("Error: Failed to generate a token")
		log.Println(err)
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
		account.Id,
	)
	if err != nil {
		log.Println("Error: Failed to create 'token' row")
		log.Println(err)
		return model.Token{}, err
	}

	tokenModel.AccountId = account.Id
	tokenModel.Token = token
	return tokenModel, nil
}
