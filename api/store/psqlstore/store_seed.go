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

	seedRolePermissionLink("Staff", "canReadAccount"),

	seedRolePermissionLink("Staff", "canReadAvailability"),

	seedRolePermissionLink("Staff", "canReadQualification"),

	seedRolePermissionLink("Staff", "canReadProductCategory"),

	seedRolePermissionLink("Staff", "canReadProductCategory"),

	seedRolePermissionLink("Staff", "canReadProduct"),

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

	// qualifications
	seedQualification("Qualification 1"),
	seedQualification("Qualification 2"),
	seedQualification("Qualification 3"),
	seedQualification("Qualification 4"),

	// account qualification links
	seedAccountQualificationLink(1, 1),
	seedAccountQualificationLink(1, 2),
	seedAccountQualificationLink(1, 3),
	seedAccountQualificationLink(1, 4),

	seedAccountQualificationLink(2, 1),
	seedAccountQualificationLink(2, 2),

	// availabilities
	seedAvailability(1, "2022-05-10T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),
	seedAvailability(1, "2022-05-11T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),
	seedAvailability(1, "2022-05-12T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),

	seedAvailability(2, "2022-05-10T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),
	seedAvailability(2, "2022-05-11T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),
	seedAvailability(2, "2022-05-12T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),
	seedAvailability(2, "2022-05-13T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),
	seedAvailability(2, "2022-05-14T09:00:00.00Z", "2022-05-10T17:00:00.00Z"),

	// product categories
	seedProductCategory("Product Category 1"),
	seedProductCategory("Product Category 2"),
	seedProductCategory("Product Category 3"),

	// products
	seedProduct(
		1,
		"Product 1",
		"Product 1 description.",
		14.99,
		2.50,
		1.5,
	),
	seedProduct(
		2,
		"Product 2",
		"Product 2 description.",
		24.99,
		4.50,
		2.5,
	),
	seedProduct(
		2,
		"Product 3",
		"Product 3 description.",
		34.99,
		6.50,
		3.5,
	),
	seedProduct(
		3,
		"Product 4",
		"Product 4 description.",
		44.99,
		8.50,
		4.5,
	),

	// product qualification link
	seedProductQualificationLink("Product 1", "Qualification 1"),
	seedProductQualificationLink("Product 1", "Qualification 3"),
	seedProductQualificationLink("Product 2", "Qualification 1"),
	seedProductQualificationLink("Product 2", "Qualification 2"),

	// booking
	seedBooking(
		1,
		1,
		"stripe_id",
		"Customer 1",
		"customer@email.com",
		"07123456789",
		"2022-05-10T09:00:00.00Z",
		3.5,
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

func seedBooking(
	productId int,
	accountId int,
	customerStripeId string,
	customerName string,
	customerEmail string,
	customerMobile string,
	date string,
	duration float64,
) string {
	return fmt.Sprintf(`
		INSERT INTO booking (
			product_id,
			account_id,
			customer_stripe_id,
			customer_name,
			customer_email,
			customer_mobile,
			date,
			duration
		)
		SELECT
			product_id,
			account_id,
			customer_stripe_id,
			customer_name,
			customer_email,
			customer_mobile,
			date,
			duration
		FROM
			booking
		UNION
		VALUES (
			%d,
			%d,
			'%s',
			'%s',
			'%s',
			'%s',
			'%s'::timestamptz,
			%f
		)
		EXCEPT
		SELECT
			product_id,
			account_id,
			customer_stripe_id,
			customer_name,
			customer_email,
			customer_mobile,
			date,
			duration
		FROM
			booking
		;`,
		productId,
		accountId,
		customerStripeId,
		customerName,
		customerEmail,
		customerMobile,
		date,
		duration,
	)
}

// Executes each seed query in succession from start to finish, populating the
// database with any seed data necessary.
func (s *PsqlStore) GenerateSeedData() {
	for _, seed := range seeds {
		s.Exec(seed)
	}
}
