package psqlstore

import (
	"fmt"
	"log"
	"strings"

	"salon-booking-guru/store"
	"salon-booking-guru/store/model"

	"golang.org/x/crypto/bcrypt"
)

// PsqlAuthenticateStore receives a pointer to an PsqlStore.
type PsqlAuthenticateStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlAuthenticateStore.
func (s *PsqlStore) Authenticate() store.AuthenticateStore {
	return &PsqlAuthenticateStore{s}
}

// Authenticates a set of credentials.
//
// Returns a AuthenticateResponse struct, and any errors encountered.
func (s *PsqlAuthenticateStore) AuthenticateCredentials(credentials model.Credentials) (model.AuthenticateResponse, error) {
	var accountId int
	var hashedPassword string

	err := s.db.QueryRow(`
		SELECT
			id,
			password
		FROM
			account
		WHERE
			email = $1
		;`,
		strings.ToLower(credentials.Email),
	).Scan(
		&accountId,
		&hashedPassword,
	)
	if err != nil {
		log.Println("Error: Failed to find account row with matching email")
		log.Println(err)
		return model.AuthenticateResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(credentials.Password),
	)
	if err != nil {
		log.Println(fmt.Sprintf("Error: Failed login attempt by '%s'", credentials.Email))
		log.Println(err)
		return model.AuthenticateResponse{}, err
	}

	token, err := s.Token().Generate(accountId)
	if err != nil {
		return model.AuthenticateResponse{}, err
	}

	accountInfo, err := s.Account().GetInfo(accountId)
	if err != nil {
		return model.AuthenticateResponse{}, err
	}

	return model.AuthenticateResponse{
		AccountInfo: accountInfo,
		Token:       token,
	}, nil
}
