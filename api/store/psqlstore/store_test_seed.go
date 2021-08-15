package psqlstore

var testSeeds = []string{
	// roles
	seedRole("Administrator"),
	seedRole("Owner"),
	seedRole("Staff"),
	seedRole("PermissionTest"),

	// permissions
	seedPermission("canReadAccount"),
	seedPermission("canCreateAccount"),
	seedPermission("canUpdateAccount"),
	seedPermission("canDeleteAccount"),

	seedPermission("canReadQualification"),
	seedPermission("canCreateQualification"),
	seedPermission("canUpdateQualification"),
	seedPermission("canDeleteQualification"),

	seedPermission("canPermissionTest"),

	// role permission links
	seedRolePermissionLink("Administrator", "canReadAccount"),
	seedRolePermissionLink("Administrator", "canCreateAccount"),
	seedRolePermissionLink("Administrator", "canUpdateAccount"),
	seedRolePermissionLink("Administrator", "canDeleteAccount"),
	seedRolePermissionLink("Administrator", "canReadQualification"),
	seedRolePermissionLink("Administrator", "canCreateQualification"),
	seedRolePermissionLink("Administrator", "canUpdateQualification"),
	seedRolePermissionLink("Administrator", "canDeleteQualification"),

	seedRolePermissionLink("Owner", "canReadAccount"),
	seedRolePermissionLink("Owner", "canCreateAccount"),
	seedRolePermissionLink("Owner", "canUpdateAccount"),
	seedRolePermissionLink("Owner", "canDeleteAccount"),
	seedRolePermissionLink("Owner", "canReadQualification"),
	seedRolePermissionLink("Owner", "canCreateQualification"),
	seedRolePermissionLink("Owner", "canUpdateQualification"),
	seedRolePermissionLink("Owner", "canDeleteQualification"),

	seedRolePermissionLink("Staff", "canReadAccount"),
	seedRolePermissionLink("Staff", "canReadQualification"),

	seedRolePermissionLink("PermissionTest", "canPermissionTest"),

	// accounts
	seedAccount(
		"admin@example.com",
		"Adam",
		"Appleby",
		// This is a bcrypt hash for the string "password"
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"Administrator",
	),
	seedAccount(
		"owner@example.com",
		"Beatrice",
		"Brown",
		// This is a bcrypt hash for the string "password"
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"Owner",
	),
	seedAccount(
		"staff1@example.com",
		"Cameron",
		"Callaway",
		// This is a bcrypt hash for the string "password"
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"Staff",
	),
	seedAccount(
		"staff2@example.com",
		"Dahlia",
		"Davidson",
		// This is a bcrypt hash for the string "password"
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"Staff",
	),
	seedAccount(
		"permissiontest@example.com",
		"Edgar",
		"Evans",
		// This is a bcrypt hash for the string "password"
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"PermissionTest",
	),

	// qualifications
	seedQualification("Qualification 1"),
	seedQualification("Qualification 2"),
	seedQualification("Qualification 3"),
}

// Executes each seed query in succession from start to finish, populating the
// database with any seed data necessary.
func (s *PsqlStore) GenerateTestSeedData() {
	for _, seed := range testSeeds {
		s.Exec(seed)
	}
}
