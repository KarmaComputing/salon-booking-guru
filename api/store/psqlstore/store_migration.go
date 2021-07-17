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
			role_id int NOT NULL,
			first_name varchar(64) NOT NULL,
			last_name varchar(64) NOT NULL,
			email varchar(254) NOT NULL,
			password varchar(64) NOT NULL,
			mobile_number varchar(32),
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
	Migration{
		`CREATE TABLE IF NOT EXISTS availability (
			id serial,
			account_id int NOT NULL,
			start_date TIMESTAMP NOT NULL,
			end_date TIMESTAMP NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (account_id) REFERENCES account(id)
		);`,
		`DROP TABLE IF EXISTS availability;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS qualification (
			id serial,
			name varchar NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		);`,
		`DROP TABLE IF EXISTS qualification;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS account_qualification_link (
			id serial,
			account_id int NOT NULL,
			qualification_id int NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (account_id) REFERENCES account(id),
			FOREIGN KEY (qualification_id) REFERENCES qualification(id)
		);`,
		`DROP TABLE IF EXISTS account_qualification_link;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS product_category (
			id serial,
			name varchar NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		);`,
		`DROP TABLE IF EXISTS product_category;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS product (
			id serial,
			product_category_id int,
			name varchar NOT NULL,
			description varchar,
			price real NOT NULL DEFAULT 0,
			deposit real NOT NULL DEFAULT 0,
			duration real,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (product_category_id) REFERENCES product_category(id)
		);`,
		`DROP TABLE IF EXISTS product;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS product_qualification_link (
			id serial,
			product_id int NOT NULL,
			qualification_id int NOT NULL,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (product_id) REFERENCES product(id),
			FOREIGN KEY (qualification_id) REFERENCES qualification(id)
		);`,
		`DROP TABLE IF EXISTS product_qualification_link;`,
	},
	Migration{
		`CREATE TABLE IF NOT EXISTS booking (
			id serial,
			product_id int NOT NULL,
			account_id int NOT NULL,
			customer_stripe_id varchar NOT NULL,
			customer_name varchar NOT NULL,
			customer_email varchar NOT NULL,
			customer_mobile varchar NOT NULL,
			date TIMESTAMP,
			duration real,
			create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modify_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (product_id) REFERENCES product(id),
			FOREIGN KEY (account_id) REFERENCES account(id)
		);`,
		`DROP TABLE IF EXISTS booking;`,
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
