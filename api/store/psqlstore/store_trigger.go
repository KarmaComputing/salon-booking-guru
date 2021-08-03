package psqlstore

var tableNames = []string{
	"role",
	"permission",
	"role_permission_link",
	"account",
	"token",
	"availability",
	"qualification",
	"account_qualification_link",
	"product_category",
	"product",
	"product_qualification_link",
	"booking",
}

// Executes each function definition query in succession from start to finish.
func (s *PsqlStore) InitTriggers() {
	for _, tableName := range tableNames {
		s.Exec(
			`
			DROP TRIGGER IF EXISTS
				sync_lastmod
			ON
				` + tableName + `
			;

			CREATE TRIGGER
				sync_lastmod
			BEFORE UPDATE ON
				` + tableName + `
			FOR EACH ROW EXECUTE PROCEDURE
				sync_lastmod()
			;`,
		)
	}
}
