package psqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"salon-booking-guru/store"
)

// assert store implementation
var _ store.Store = &PsqlStore{}

// PsqlStore receives a pointer to an sql.DB
type PsqlStore struct {
	db *sql.DB
}

// Open a connection to a psql database and return a pointer to a PsqlStore
// struct with a pointer to the db connection.
func Open() (*PsqlStore, error) {
	host := os.Getenv("SBG_DB_HOST")
	port := os.Getenv("SBG_DB_PORT")
	user := os.Getenv("SBG_DB_USER")
	password := os.Getenv("SBG_DB_PASSWORD")
	dbname := os.Getenv("SBG_DB_DBNAME")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	var s PsqlStore
	var err error
	s.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = TestDatabase(s)
	if err != nil {
		log.Println("Shutting down")
		os.Exit(1)
	}

	s.Up()
	s.DefineFunctions()
	s.InitTriggers()

	return &s, nil
}

// Open a connection to a psql database and return a pointer to a PsqlStore
// struct with a pointer to the db connection.
func OpenTest() (*PsqlStore, error) {
	host := os.Getenv("SBG_DB_HOST")
	port := os.Getenv("SBG_DB_PORT")
	user := os.Getenv("SBG_DB_USER")
	password := os.Getenv("SBG_DB_PASSWORD")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		host,
		port,
		user,
		password,
	)

	var s PsqlStore
	var err error
	s.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = s.db.Exec(`
		SELECT
			pg_terminate_backend(pg_stat_activity.pid)
		FROM
			pg_stat_activity
		WHERE
			pg_stat_activity.datname = 'salon_booking_guru_test'
		AND
			pid <> pg_backend_pid()
		;`,
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_, err = s.db.Exec("DROP DATABASE IF EXISTS salon_booking_guru_test;")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_, err = s.db.Exec("CREATE DATABASE salon_booking_guru_test;")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	s.db.Close()

	connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=salon_booking_guru_test sslmode=disable",
		host,
		port,
		user,
		password,
	)

	s.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = TestDatabase(s)
	if err != nil {
		log.Println("Shutting down")
		os.Exit(1)
	}

	s.Up()
	s.DefineFunctions()
	s.InitTriggers()

	return &s, nil
}

// Test connection to the database
func TestDatabase(s PsqlStore) error {
	var err error
	counter := 0
	interval := 10
	attempts := 10

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()
	for ; true; <-ticker.C {
		counter++
		err = s.db.Ping()
		if err != nil {
			if counter > attempts {
				err = errors.New(
					"Failed to connect to the database",
				)
				log.Println(err)
				return err
			}
			log.Println(err)
			log.Println(
				fmt.Sprintf(
					"Will attempt to retry in %ds [Attempt %d/%d]",
					interval,
					counter,
					attempts,
				),
			)
			continue
		}
		break
	}
	return nil
}

// Close the db connection.
func (s *PsqlStore) Close() {
	s.Close()
}

// Execute a single query using the db connection.
func (s *PsqlStore) Exec(query string) {
	if _, err := s.db.Exec(query); err != nil {
		log.Println(err)
		fmt.Println(query)
	}
}
