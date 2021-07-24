package psqlstore

var testSeeds = []string{
	// roles
	seedRole("Administrator"),
	seedRole("Owner"),
	seedRole("Staff"),

	// permissions
	seedPermission("Administrator"),
	seedPermission("Owner"),
	seedPermission("Staff"),

	// role permission links
	seedRolePermissionLink("Administrator", "Administrator"),
	seedRolePermissionLink("Owner", "Owner"),
	seedRolePermissionLink("Staff", "Staff"),

	// accounts
	seedAccount(
		"admin@example.com",
		"Adam",
		"Appleby",
		"$2y$10$FdhRrvtzETtcFsezkFlX.ujOc9H6v3LnmOd8ITZ7mWPIjRIEvgDa.",
		"Administrator",
	),
	seedAccount(
		"owner@example.com",
		"Beatrice",
		"Brown",
		"$2y$10$FdhRrvtzETtcFsezkFlX.ujOc9H6v3LnmOd8ITZ7mWPIjRIEvgDa.",
		"Owner",
	),
	seedAccount(
		"staff1@example.com",
		"Cameron",
		"Callaway",
		"$2y$10$FdhRrvtzETtcFsezkFlX.ujOc9H6v3LnmOd8ITZ7mWPIjRIEvgDa.",
		"Staff",
	),
	seedAccount(
		"staff2@example.com",
		"Dahlia",
		"Davidson",
		"$2y$10$FdhRrvtzETtcFsezkFlX.ujOc9H6v3LnmOd8ITZ7mWPIjRIEvgDa.",
		"Staff",
	),
}

// Executes each seed query in succession from start to finish, populating the
// database with any seed data necessary.
func (s *PsqlStore) GenerateTestSeedData() {
	for _, seed := range testSeeds {
		s.Exec(seed)
	}
}
