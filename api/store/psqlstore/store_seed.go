package psqlstore

import "fmt"

var seeds = []string{
	// roles
	seedRole("Administrator"),
	seedRole("Owner"),
	seedRole("Staff"),

	// permissions
	seedPermission("canReadAccount"),
	seedPermission("canCreateAccount"),
	seedPermission("canUpdateAccount"),
	seedPermission("canDeleteAccount"),

	// role permission links
	seedRolePermissionLink("Administrator", "canReadAccount"),
	seedRolePermissionLink("Administrator", "canCreateAccount"),
	seedRolePermissionLink("Administrator", "canUpdateAccount"),
	seedRolePermissionLink("Administrator", "canDeleteAccount"),

	seedRolePermissionLink("Owner", "canReadAccount"),
	seedRolePermissionLink("Owner", "canCreateAccount"),
	seedRolePermissionLink("Owner", "canUpdateAccount"),
	seedRolePermissionLink("Owner", "canDeleteAccount"),

	seedRolePermissionLink("Staff", "canReadAccount"),

	// accounts
	seedAccount(
		"jamie@scollay.uk",
		"Jamie",
		"Scollay",
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"Administrator",
	),
	seedAccount(
		"changeme@owner.com",
		"Owner",
		"Name",
		"$2y$10$tIU8Z5tQXN7oBoeG24hzQuucWjVHyw/6UuDUtA88Ae8rlIhA.hc7e",
		"Owner",
	),
}

func seedRole(
	name string,
) string {
	return fmt.Sprintf(`
		INSERT INTO role (name)
		SELECT name FROM role
		UNION
		VALUES ('%s')
		EXCEPT
		SELECT name FROM role
		;`,
		name,
	)
}

func seedPermission(name string) string {
	return fmt.Sprintf(`
		INSERT INTO permission (name)
		SELECT name FROM permission
		UNION
		VALUES ('%s')
		EXCEPT
		SELECT name FROM permission
		;`,
		name,
	)
}

func seedRolePermissionLink(roleName string, permissionName string) string {
	return fmt.Sprintf(`
		INSERT INTO role_permission_link (role_id, permission_id)
		SELECT role_id, permission_id FROM role_permission_link
		UNION
		VALUES (
			(SELECT id FROM role WHERE name = '%s'),
			(SELECT id FROM permission WHERE name = '%s')
		)
		EXCEPT
		SELECT role_id, permission_id FROM role_permission_link;
		`,
		roleName,
		permissionName,
	)
}

func seedAccount(
	email string,
	firstName string,
	lastName string,
	password string,
	roleName string,
) string {
	return fmt.Sprintf(`
		INSERT INTO account (
			email, first_name, last_name, password, role_id
		)
		SELECT
			email, first_name, last_name, password, role_id
		FROM
			account
		UNION
		VALUES (
			'%s',
			'%s',
			'%s',
			'%s',
			(SELECT id FROM role WHERE name = '%s')
		)
		EXCEPT
		SELECT
			email, first_name, last_name, password, role_id
		FROM
			account;
		`,
		email,
		firstName,
		lastName,
		password,
		roleName,
	)
}

// Executes each seed query in succession from start to finish, populating the
// database with any seed data necessary.
func (s *PsqlStore) GenerateSeedData() {
	for _, seed := range seeds {
		s.Exec(seed)
	}
}
