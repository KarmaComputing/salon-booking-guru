package psqlstore

var functions = []string{
	`
	CREATE OR REPLACE FUNCTION sync_lastmod() RETURNS trigger AS $$
	BEGIN
		NEW.modify_date := NOW();
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	`,
}

// Executes each function definition query in succession from start to finish.
func (s *PsqlStore) DefineFunctions() {
	for _, function := range functions {
		s.Exec(function)
	}
}
