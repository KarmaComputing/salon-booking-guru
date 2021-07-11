package psqlstore

// A single Migration has an 'Up' query and a corresponding 'Down' query, the
// 'Down' query should be the exact inverse of the 'Up' query, essentially
// reversing any changes made.
type Migration struct {
	Up   string
	Down string
}

var migrations = []Migration{
	Migration{
		`CREATE TABLE IF NOT EXISTS role (
			id serial,
			name varchar(50) NOT NULL DEFAULT '',
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		);`,
		`DROP TABLE IF EXISTS role;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS permission (
			id serial,
			name varchar(50) NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		);`,
		`DROP TABLE IF EXISTS permission;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS role_permission_link (
			id serial,
			role_id int NOT NULL,
			permission_id int NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (role_id) REFERENCES role(id),
			FOREIGN KEY (permission_id) REFERENCES permission(id)
		);`,
		`DROP TABLE IF EXISTS role_permission_link;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS account (
			id serial,
			email varchar(254) NOT NULL,
			first_name varchar(64) NOT NULL,
			last_name varchar(64) NOT NULL,
			password varchar(64) NOT NULL,
			role_id int NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (role_id) REFERENCES role(id)
		);`,
		`DROP TABLE IF EXISTS account;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS token (
			id serial,
			account_id int NOT NULL,
			token varchar(100) NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (account_id) REFERENCES account(id)
		);`,
		`DROP TABLE IF EXISTS token;`,
	},
}

// Execute each migration.Up query in succession from start to finish.
func (s *PsqlStore) Up() {
	for _, migration := range migrations {
		s.Exec(migration.Up)
	}
	s.GenerateSeedData()
}

// Execute each migration.Down query in succession from start to finish.
func (s *PsqlStore) Down() {
	for i, j := 0, len(migrations)-1; i < j; i, j = i+1, j-1 {
		migrations[i], migrations[j] = migrations[j], migrations[i]
	}
	for _, migration := range migrations {
		s.Exec(migration.Down)
	}
}
