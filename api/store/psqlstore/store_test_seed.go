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

	seedPermission("canReadAvailability"),
	seedPermission("canCreateAvailability"),
	seedPermission("canUpdateAvailability"),
	seedPermission("canDeleteAvailability"),

	seedPermission("canReadQualification"),
	seedPermission("canCreateQualification"),
	seedPermission("canUpdateQualification"),
	seedPermission("canDeleteQualification"),

	seedPermission("canReadProductCategory"),
	seedPermission("canCreateProductCategory"),
	seedPermission("canUpdateProductCategory"),
	seedPermission("canDeleteProductCategory"),

	seedPermission("canPermissionTest"),

	// role permission links
	seedRolePermissionLink("Administrator", "canReadAccount"),
	seedRolePermissionLink("Administrator", "canCreateAccount"),
	seedRolePermissionLink("Administrator", "canUpdateAccount"),
	seedRolePermissionLink("Administrator", "canDeleteAccount"),

	seedRolePermissionLink("Administrator", "canReadAvailability"),
	seedRolePermissionLink("Administrator", "canCreateAvailability"),
	seedRolePermissionLink("Administrator", "canUpdateAvailability"),
	seedRolePermissionLink("Administrator", "canDeleteAvailability"),

	seedRolePermissionLink("Administrator", "canReadQualification"),
	seedRolePermissionLink("Administrator", "canCreateQualification"),
	seedRolePermissionLink("Administrator", "canUpdateQualification"),
	seedRolePermissionLink("Administrator", "canDeleteQualification"),

	seedRolePermissionLink("Administrator", "canReadProductCategory"),
	seedRolePermissionLink("Administrator", "canCreateProductCategory"),
	seedRolePermissionLink("Administrator", "canUpdateProductCategory"),
	seedRolePermissionLink("Administrator", "canDeleteProductCategory"),

	seedRolePermissionLink("Owner", "canReadAccount"),
	seedRolePermissionLink("Owner", "canCreateAccount"),
	seedRolePermissionLink("Owner", "canUpdateAccount"),
	seedRolePermissionLink("Owner", "canDeleteAccount"),

	seedRolePermissionLink("Owner", "canReadAvailability"),
	seedRolePermissionLink("Owner", "canCreateAvailability"),
	seedRolePermissionLink("Owner", "canUpdateAvailability"),
	seedRolePermissionLink("Owner", "canDeleteAvailability"),

	seedRolePermissionLink("Owner", "canReadQualification"),
	seedRolePermissionLink("Owner", "canCreateQualification"),
	seedRolePermissionLink("Owner", "canUpdateQualification"),
	seedRolePermissionLink("Owner", "canDeleteQualification"),

	seedRolePermissionLink("Owner", "canReadProductCategory"),
	seedRolePermissionLink("Owner", "canCreateProductCategory"),
	seedRolePermissionLink("Owner", "canUpdateProductCategory"),
	seedRolePermissionLink("Owner", "canDeleteProductCategory"),

	seedRolePermissionLink("Staff", "canReadAccount"),

	seedRolePermissionLink("Staff", "canReadAvailability"),

	seedRolePermissionLink("Staff", "canReadQualification"),

	seedRolePermissionLink("Staff", "canReadProductCategory"),

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
	seedQualification("Qualification 4"),

	// account qualification links
	seedAccountQualificationLink(2, 2),
	seedAccountQualificationLink(2, 3),
	seedAccountQualificationLink(2, 4),

	// availabilities
	seedAvailability(3, "2021-05-10T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),
	seedAvailability(3, "2021-05-11T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),
	seedAvailability(3, "2021-05-12T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),

	seedAvailability(4, "2021-05-10T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),
	seedAvailability(4, "2021-05-11T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),
	seedAvailability(4, "2021-05-12T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),
	seedAvailability(4, "2021-05-13T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),
	seedAvailability(4, "2021-05-14T09:00:00.00Z", "2021-05-10T17:00:00.00Z"),

	// product categories
	seedProductCategory("Product Category 1"),
	seedProductCategory("Product Category 2"),
	seedProductCategory("Product Category 3"),
}

// Executes each seed query in succession from start to finish, populating the
// database with any seed data necessary.
func (s *PsqlStore) GenerateTestSeedData() {
	for _, seed := range testSeeds {
		s.Exec(seed)
	}
}
