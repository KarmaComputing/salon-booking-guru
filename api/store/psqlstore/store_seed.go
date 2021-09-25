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

	seedPermission("canReadProduct"),
	seedPermission("canCreateProduct"),
	seedPermission("canUpdateProduct"),
	seedPermission("canDeleteProduct"),

	seedPermission("canReadRole"),

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

	seedRolePermissionLink("Administrator", "canReadProduct"),
	seedRolePermissionLink("Administrator", "canCreateProduct"),
	seedRolePermissionLink("Administrator", "canUpdateProduct"),
	seedRolePermissionLink("Administrator", "canDeleteProduct"),

	seedRolePermissionLink("Administrator", "canReadRole"),

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

	seedRolePermissionLink("Owner", "canReadProduct"),
	seedRolePermissionLink("Owner", "canCreateProduct"),
	seedRolePermissionLink("Owner", "canUpdateProduct"),
	seedRolePermissionLink("Owner", "canDeleteProduct"),

	seedRolePermissionLink("Owner", "canReadRole"),

	seedRolePermissionLink("Staff", "canReadAccount"),

	seedRolePermissionLink("Staff", "canReadAvailability"),

	seedRolePermissionLink("Staff", "canReadQualification"),

	seedRolePermissionLink("Staff", "canReadProductCategory"),

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

	seedAccount(
		"changemehello@owner.com",
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

func seedQualification(
	name string,
) string {
	return fmt.Sprintf(`
		INSERT INTO qualification (
			name
		)
		SELECT
			name
		FROM
			qualification
		UNION
		VALUES (
			'%s'
		)
		EXCEPT
		SELECT
			name
		FROM
			qualification
		;`,
		name,
	)
}

func seedAccountQualificationLink(accountId int, qualificationId int) string {
	return fmt.Sprintf(`
		INSERT INTO account_qualification_link (account_id, qualification_id)
		SELECT account_id, qualification_id FROM account_qualification_link
		UNION
		VALUES (
			%d,
			%d
		)
		EXCEPT
		SELECT account_id, qualification_id FROM account_qualification_link;
		`,
		accountId,
		qualificationId,
	)
}

func seedAvailability(
	accountId int,
	startDate string,
	endDate string,
) string {
	return fmt.Sprintf(`
		INSERT INTO availability (
			account_id,
			start_date,
			end_date
		)
		SELECT
			account_id,
			start_date,
			end_date
		FROM
			availability
		UNION
		VALUES (
			%d,
			'%s'::timestamptz,
			'%s'::timestamptz
		)
		EXCEPT
		SELECT
			account_id,
			start_date,
			end_date
		FROM
			availability
		;`,
		accountId,
		startDate,
		endDate,
	)
}

func seedProductCategory(
	name string,
) string {
	return fmt.Sprintf(`
		INSERT INTO product_category (
			name
		)
		SELECT
			name
		FROM
			product_category
		UNION
		VALUES (
			'%s'
		)
		EXCEPT
		SELECT
			name
		FROM
			product_category
		;`,
		name,
	)
}

func seedProduct(
	productCategoryId int,
	name string,
	description string,
	price float64,
	deposit float64,
	duration float64,
) string {
	return fmt.Sprintf(`
		INSERT INTO product (
			product_category_id,
			name,
			description,
			price,
			deposit,
			duration
		)
		SELECT
			product_category_id,
			name,
			description,
			price,
			deposit,
			duration
		FROM
			product
		UNION
		VALUES (
			%d,
			'%s',
			'%s',
			%f,
			%f,
			%f
		)
		EXCEPT
		SELECT
			product_category_id,
			name,
			description,
			price,
			deposit,
			duration
		FROM
			product
		;`,
		productCategoryId,
		name,
		description,
		price,
		deposit,
		duration,
	)
}

func seedProductQualificationLink(productName string, qualificationName string) string {
	return fmt.Sprintf(`
		INSERT INTO product_qualification_link (product_id, qualification_id)
		SELECT product_id, qualification_id FROM product_qualification_link
		UNION
		VALUES (
			(SELECT id FROM product WHERE name = '%s'),
			(SELECT id FROM qualification WHERE name = '%s')
		)
		EXCEPT
		SELECT product_id, qualification_id FROM product_qualification_link;
		`,
		productName,
		qualificationName,
	)
}

// Executes each seed query in succession from start to finish, populating the
// database with any seed data necessary.
func (s *PsqlStore) GenerateSeedData() {
	for _, seed := range seeds {
		s.Exec(seed)
	}
}
