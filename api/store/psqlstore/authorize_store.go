package psqlstore

import (
	"errors"
	"log"

	"salon-booking-guru/store"

	"github.com/lib/pq"
)

// PsqlAuthorizeStore receives a pointer to an PsqlStore.
type PsqlAuthorizeStore struct {
	*PsqlStore
}

// Returns the a pointer to a PsqlAuthorizeStore.
func (s *PsqlStore) Authorize() store.AuthorizeStore {
	return &PsqlAuthorizeStore{s}
}

// Authorizes a bearer token by checking it exists in the database, and that the
// associated account has the required permissions, if no required permissions
// are specified it will simply check the token is found.
//
// Returns any errors encountered.
func (s *PsqlAuthorizeStore) AuthorizeToken(bearerToken string, requiredPermissions []string) error {
	if len(requiredPermissions) == 0 {
		var accountId int
		err := s.db.QueryRow(`
			SELECT
				account.id
			FROM
				account
			INNER JOIN
				token
			ON
				account.id = token.account_id
			WHERE
				'Bearer ' || token.token = $1
			;`,
			bearerToken,
		).Scan(
			&accountId,
		)
		if err != nil {
			return err
		}
		return nil
	}

	isAuthorized := false
	err := s.db.QueryRow(`
		SELECT ARRAY(
			SELECT
				permission.name
			FROM
				account
			INNER JOIN
				token
			ON
				account.id = token.account_id
			INNER JOIN
				role
			ON
				account.role_id = role.id
			INNER JOIN
				role_permission_link
			ON
				role.id = role_permission_link.role_id
			INNER JOIN
				permission
			ON
				role_permission_link.permission_id = permission.id
			WHERE
				'Bearer ' || token.token = $1
		) && $2
		;`,
		bearerToken,
		pq.Array(requiredPermissions),
	).Scan(
		&isAuthorized,
	)
	if err != nil {
		log.Println("Error: Failed to execute authorization query")
		log.Println(err)
		return err
	}

	if !isAuthorized {
		return errors.New("Error: Failed to authorize token")
	}

	return nil
}
