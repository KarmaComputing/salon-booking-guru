package psqlstore

import (
	"database/sql"
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
	var s PsqlStore
	go InitDatabase(&s)
	return &s, nil
}

// Open a connection to a psql database and return a pointer to a PsqlStore
// struct with a pointer to the db connection.
func OpenTest() (*PsqlStore, error) {
	var s PsqlStore
	err := InitTestDatabase(&s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// Initialise the database
func InitDatabase(s *PsqlStore) {
	host := os.Getenv("SALON_BOOKING_GURU_DB_HOST")
	port := os.Getenv("SALON_BOOKING_GURU_DB_PORT")
	user := os.Getenv("SALON_BOOKING_GURU_DB_USER")
	password := os.Getenv("SALON_BOOKING_GURU_DB_PASSWORD")
	dbName := os.Getenv("SALON_BOOKING_GURU_DB_DBNAME")

	waitPeriod := time.Second * 1
	maxWaitPeriod := time.Minute * 5

	for {
		if waitPeriod != time.Second*1 {
			log.Printf("Attempting to reconnect in %ds", waitPeriod/2000000000)
			time.Sleep(waitPeriod)
		}
		waitPeriod *= 2
		if waitPeriod > maxWaitPeriod {
			waitPeriod = maxWaitPeriod
		}

		connectionString := fmt.Sprintf(
			"host=%s port=%s user=%s password='%s' sslmode=disable",
			host,
			port,
			user,
			password,
		)

		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Println(err)
			continue
		}
		s.db = db
		defer s.db.Close()

		_, err = s.db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if err != nil &&
			err.Error() != fmt.Sprintf("pq: database \"%s\" already exists", dbName) {
			log.Println(err)
			continue
		}

		connectionString = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			password,
			dbName,
		)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Println(err)
			continue
		}
		s.db = db

		err = s.db.Ping()
		if err != nil {
			log.Println(err)
			continue
		}

		s.Up()
		s.GenerateSeedData()
		s.DefineFunctions()
		s.InitTriggers()

		log.Println("Connection to the database successfully established")
		break
	}
}

// Initialise the test database
func InitTestDatabase(s *PsqlStore) error {
	host := os.Getenv("SALON_BOOKING_GURU_DB_HOST")
	port := os.Getenv("SALON_BOOKING_GURU_DB_PORT")
	user := os.Getenv("SALON_BOOKING_GURU_DB_USER")
	password := os.Getenv("SALON_BOOKING_GURU_DB_PASSWORD")
	dbName := "salon_booking_guru_test"

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password='%s' sslmode=disable",
		host,
		port,
		user,
		password,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		return err
	}
	s.db = db
	defer s.db.Close()

	_, err = s.db.Exec(`
			SELECT
				pg_terminate_backend(pg_stat_activity.pid)
			FROM
				pg_stat_activity
			WHERE
				pg_stat_activity.datname = $1
			AND
				pid <> pg_backend_pid()
			;`,
		dbName,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = s.db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = s.db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if err != nil {
		log.Println(err)
		return err
	}

	connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		return err
	}
	s.db = db

	err = s.db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}

	s.Up()
	s.GenerateTestSeedData()
	s.DefineFunctions()
	s.InitTriggers()

	return nil
}

// Execute a single query using the db connection.
func (s *PsqlStore) Exec(query string) {
	if _, err := s.db.Exec(query); err != nil {
		log.Println(err)
		fmt.Println(query)
	}
}
